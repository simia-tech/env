package env

import (
	"os"
	"strings"
)

const separator = ","

// StringsField implements a strings field.
type StringsField struct {
	name         string
	defaultValue []string
	options      *options
}

// Strings registers a field of the provided name.
func Strings(name string, defaultValue []string, opts ...Option) *StringsField {
	field := &StringsField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Strings field."}, opts),
	}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (sf *StringsField) Name() string {
	return sf.name
}

// Value returns the field's value.
func (sf *StringsField) Value() string {
	return strings.Join(sf.Get(), separator)
}

// DefaultValue returns the field's default value.
func (sf *StringsField) DefaultValue() string {
	return strings.Join(sf.defaultValue, separator)
}

// Description returns the field's description.
func (sf *StringsField) Description() string {
	return sf.options.description(sf.DefaultValue())
}

// Get returns the field value or the default value.
func (sf *StringsField) Get() []string {
	text := os.Getenv(sf.Name())
	if text == "" {
		if sf.options.required {
			requiredError(sf)
		}
		return sf.defaultValue
	}
	values := strings.Split(text, separator)
	for _, value := range values {
		if !sf.options.isAllowedValue(value) {
			unallowedError(sf, value, sf.options.allowedValues)
			return sf.defaultValue
		}
	}
	return values
}