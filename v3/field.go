package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrMissingValue = errors.New("missing value")
	ErrInvalidValue = errors.New("invalid value")
)

type FieldType interface {
	string | int | bool
}

type Field[T FieldType] struct {
	name         string
	defaultValue T
	options      options
}

func NewField[T FieldType](name string, defaultValue T, opts ...Option) *Field[T] {
	return &Field[T]{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions(opts),
	}
}

func (f *Field[T]) Name() string {
	return f.name
}

func (f *Field[T]) Get() (T, error) {
	text, ok := os.LookupEnv(f.name)
	if !ok {
		if f.options.required {
			return f.defaultValue, fmt.Errorf("field [%s]: %w", f.name, ErrMissingValue)
		}
		return f.defaultValue, nil
	}
	text = strings.TrimSpace(text)

	if !f.options.isAllowedValue(text) {
		return f.defaultValue, fmt.Errorf("field [%s]: value [%s]: %w", f.name, text, ErrInvalidValue)
	}

	result := any(nil)

	switch any(f.defaultValue).(type) {
	case string:
		result = text

	case int:
		value, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return f.defaultValue, fmt.Errorf("field [%s]: parse int [%s]: %w", f.name, text, ErrInvalidValue)
		}
		result = int(value)

	case bool:
		switch text {
		case "1", "true", "yes":
			result = true
		case "0", "false", "no":
			result = false
		default:
			return f.defaultValue, fmt.Errorf("field [%s]: parse bool [%s]: %w", f.name, text, ErrInvalidValue)
		}
	}

	return result.(T), nil
}
