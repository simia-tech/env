package env

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/simia-tech/env/v3/internal/parser"
)

var (
	ErrMissingValue = errors.New("missing value")
	ErrInvalidValue = errors.New("invalid value")
)

type FieldItemType interface {
	string
}

type FieldType[K, V FieldItemType] interface {
	bool | int | string | []K | map[K]V
}

type FieldValue[T FieldType[K, V], K, V FieldItemType] struct {
	name         string
	location     string
	defaultValue T
	options      options
}

var nameRegexp = regexp.MustCompile("^[A-Z0-9_]+$")

func Field[T FieldType[K, V], K, V FieldItemType](name string, defaultValue T, opts ...Option) *FieldValue[T, K, V] {
	if !nameRegexp.MatchString(name) {
		panic(fmt.Sprintf("field name [%s] must only contain capital letters, numbers or underscores", name))
	}
	_, filename, line, _ := runtime.Caller(1)

	f := &FieldValue[T, K, V]{
		name:         name,
		location:     fmt.Sprintf("%s:%d", filename, line),
		defaultValue: defaultValue,
		options:      newOptions(opts),
	}
	registerField(f)
	return f
}

func (f *FieldValue[T, K, V]) Name() string {
	return f.name
}

func (f *FieldValue[T, K, V]) Description() string {
	if f.options.description != "" {
		return f.options.description
	}
	sentences := []string{label[T, K, V]() + " field."}
	if f.options.required {
		sentences = append(sentences, "Required field.")
	}
	if f.options.allowedValues != nil {
		sentences = append(sentences, fmt.Sprintf("Allowed values are %s.", joinStringValues(f.options.allowedValues)))
	}
	sentences = append(sentences, "The default value is '"+formatValue[T, K, V](f.defaultValue)+"'.")
	sentences = append(sentences, "Defined at "+f.location+".")
	return strings.Join(sentences, " ")
}

func (f *FieldValue[T, K, V]) GetRaw() (string, error) {
	text, ok := os.LookupEnv(f.name)
	if !ok {
		if f.options.required {
			return formatValue[T, K, V](f.defaultValue), fmt.Errorf("field [%s]: %w", f.name, ErrMissingValue)
		}
		return formatValue[T, K, V](f.defaultValue), nil
	}
	text = strings.TrimSpace(text)

	if !f.options.isAllowedValue(text) {
		return formatValue[T, K, V](f.defaultValue), fmt.Errorf("field [%s]: value [%s]: %w", f.name, text, ErrInvalidValue)
	}

	return text, nil
}

func (f *FieldValue[T, K, V]) GetRawOrDefault() string {
	value, _ := f.GetRaw()
	return value
}

func (f *FieldValue[T, K, V]) Get() (T, error) {
	raw, err := f.GetRaw()
	if err != nil {
		return f.defaultValue, err
	}

	result, err := parseValue[T, K, V](raw)
	if err != nil {
		return f.defaultValue, fmt.Errorf("field [%s]: %w", f.name, err)
	}

	return result, nil
}

func label[T FieldType[K, V], K, V FieldItemType]() string {
	switch any(*new(T)).(type) {
	case bool:
		return "Boolean"
	case int:
		return "Int"
	case string:
		return "String"
	case []K:
		return label[K, string, string]() + "Array"
	case map[K]V:
		return label[K, string, string]() + label[V, string, string]() + "Map"
	default:
		return "Unknown"
	}
}

func parseValue[T FieldType[K, V], K, V FieldItemType](raw string) (T, error) {
	value := *new(T)

	result := any(nil)
	switch any(value).(type) {
	case bool:
		switch raw {
		case "1", "true", "yes":
			result = true
		case "0", "false", "no":
			result = false
		default:
			return value, fmt.Errorf("parse bool [%s]: %w", raw, ErrInvalidValue)
		}

	case int:
		v, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return value, fmt.Errorf("parse int [%s]: %w", raw, ErrInvalidValue)
		}
		result = int(v)

	case string:
		result = raw

	case []K:
		v, err := parser.ParseStrings(raw)
		if err != nil {
			return value, fmt.Errorf("parse string array [%s]: %w", raw, err)
		}
		result = v

	case map[K]V:
		m, err := parser.ParseStringMap(raw)
		if err != nil {
			return value, fmt.Errorf("parse string map [%s]: %w", raw, err)
		}
		result = m

	}

	return result.(T), nil
}

func formatValue[T FieldType[K, V], K, V FieldItemType](value T) string {
	switch t := any(value).(type) {
	case bool:
		if t {
			return "true"
		}
		return "false"

	case int:
		return strconv.FormatInt(int64(t), 10)

	case string:
		return t

	case []K:
		parts := []string{}
		for _, v := range t {
			parts = append(parts, formatValue[K, string, string](v))
		}
		return strings.Join(parts, ",")

	case map[K]V:
		m := map[string]string{}
		for k, v := range t {
			m[formatValue[K, string, string](k)] = formatValue[V, string, string](v)
		}
		return parser.FormatStringMap(m)

	default:
		return ""
	}
}
