package env

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/simia-tech/env/v3/internal/parser"
)

var (
	ErrMissingValue = errors.New("missing value")
	ErrInvalidValue = errors.New("invalid value")
)

type FieldType interface {
	bool | []byte | time.Duration | int | []int | string | []string | map[string]string
}

type F[T FieldType] struct {
	name         string
	location     string
	defaultValue T
	options      options
}

var nameRegexp = regexp.MustCompile("^[A-Z0-9_]+$")

func Field[T FieldType](name string, defaultValue T, opts ...Option) *F[T] {
	if !nameRegexp.MatchString(name) {
		panic(fmt.Sprintf("field name [%s] must only contain capital letters, numbers or underscores", name))
	}
	_, filename, line, _ := runtime.Caller(1)

	f := &F[T]{
		name:         name,
		location:     fmt.Sprintf("%s:%d", filename, line),
		defaultValue: defaultValue,
		options:      newOptions(opts),
	}
	registerField(f)
	return f
}

func (f *F[T]) Name() string {
	return f.name
}

func (f *F[T]) Description() string {
	if f.options.description != "" {
		return f.options.description
	}
	sentences := []string{label[T]() + " field."}
	if f.options.required {
		sentences = append(sentences, "Required field.")
	}
	if f.options.allowedValues != nil {
		sentences = append(sentences, fmt.Sprintf("Allowed values are %s.", joinStringValues(f.options.allowedValues)))
	}
	sentences = append(sentences, "The default value is '"+formatValue[T](f.defaultValue)+"'.")
	sentences = append(sentences, "Defined at "+f.location+".")
	return strings.Join(sentences, " ")
}

func (f *F[T]) GetRaw() (string, error) {
	text, ok := os.LookupEnv(f.name)
	if !ok {
		if f.options.required {
			return formatValue[T](f.defaultValue), fmt.Errorf("field [%s]: %w", f.name, ErrMissingValue)
		}
		return formatValue[T](f.defaultValue), nil
	}
	text = strings.TrimSpace(text)

	if !f.options.isAllowedValue(text) {
		return formatValue[T](f.defaultValue), fmt.Errorf("field [%s]: value [%s]: %w", f.name, text, ErrInvalidValue)
	}

	return text, nil
}

func (f *F[T]) GetRawOrDefault() string {
	value, _ := f.GetRaw()
	return value
}

func (f *F[T]) MustGetRaw() string {
	value, err := f.GetRaw()
	if err != nil {
		panic(err)
	}
	return value
}

func (f *F[T]) Get() (T, error) {
	raw, err := f.GetRaw()
	if err != nil {
		return f.defaultValue, err
	}

	result, err := parseValue[T](raw)
	if err != nil {
		return f.defaultValue, fmt.Errorf("field [%s]: %w", f.name, err)
	}

	return result, nil
}

func (f *F[T]) GetOrDefault() T {
	value, _ := f.Get()
	return value
}

func (f *F[T]) MustGet() T {
	value, err := f.Get()
	if err != nil {
		panic(err)
	}
	return value
}

func label[T FieldType]() string {
	switch any(*new(T)).(type) {
	case bool:
		return "Boolean"
	case []byte:
		return "Bytes"
	case time.Duration:
		return "Duration"
	case int:
		return "Int"
	case []int:
		return "IntArray"
	case string:
		return "String"
	case []string:
		return "StringArray"
	case map[string]string:
		return "StringStringMap"
	default:
		return "Unknown"
	}
}

func parseValue[T FieldType](raw string) (T, error) {
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

	case []byte:
		v, err := hex.DecodeString(raw)
		if err != nil {
			return value, fmt.Errorf("parse hex [%s]: %w", raw, ErrInvalidValue)
		}
		result = v

	case time.Duration:
		v, err := time.ParseDuration(raw)
		if err != nil {
			return value, fmt.Errorf("parse duration [%s]: %w", raw, ErrInvalidValue)
		}
		result = v

	case int:
		v, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return value, fmt.Errorf("parse int [%s]: %w", raw, ErrInvalidValue)
		}
		result = int(v)

	case []int:
		v, err := parser.ParseInts(raw)
		if err != nil {
			return value, fmt.Errorf("parse int array [%s]: %w", raw, err)
		}
		result = v

	case string:
		result = raw

	case []string:
		v, err := parser.ParseStrings(raw)
		if err != nil {
			return value, fmt.Errorf("parse string array [%s]: %w", raw, err)
		}
		result = v

	case map[string]string:
		m, err := parser.ParseStringMap(raw)
		if err != nil {
			return value, fmt.Errorf("parse string map [%s]: %w", raw, err)
		}
		result = m

	}

	return result.(T), nil
}

func formatValue[T FieldType](value T) string {
	switch t := any(value).(type) {
	case bool:
		if t {
			return "true"
		}
		return "false"

	case []byte:
		return hex.EncodeToString(t)

	case time.Duration:
		return t.String()

	case int:
		return strconv.FormatInt(int64(t), 10)

	case []int:
		return parser.FormatInts(t)

	case string:
		return t

	case []string:
		return parser.FormatStrings(t)

	case map[string]string:
		return parser.FormatStringMap(t)

	default:
		return ""
	}
}
