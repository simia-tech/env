package env

import (
	"os"
	"strconv"
)

// IntField implements a duration field.
type IntField struct {
	name         string
	defaultValue int
	options      *options
}

// Int registers a field of the provided name.
func Int(name string, defaultValue int, opts ...Option) *IntField {
	field := &IntField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Int field."}, opts),
	}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (i *IntField) Name() string {
	return i.name
}

// Value returns the field's value.
func (i *IntField) Value() string {
	return strconv.Itoa(i.Get())
}

// DefaultValue returns the field's default value.
func (i *IntField) DefaultValue() string {
	return strconv.Itoa(i.defaultValue)
}

// Description returns the field's description.
func (i *IntField) Description() string {
	return i.options.description(i.DefaultValue())
}

// Get returns the field value or the default value.
func (i *IntField) Get() int {
	text := os.Getenv(i.Name())
	if text == "" {
		if i.options.required {
			requiredError(i)
		}
		return i.defaultValue
	}
	if !i.options.isAllowedValue(text) {
		unallowedError(i, text, i.options.allowedValues)
		return i.defaultValue
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		parseError(i, "Int", text)
		return i.defaultValue
	}
	return value
}
