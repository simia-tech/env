package env

import "os"

// StringField implements a string field.
type StringField struct {
	name         string
	defaultValue string
	options      *options
}

// String registers a field of the provided name.
func String(name, defaultValue string, opts ...Option) *StringField {
	field := &StringField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"String field."}, opts),
	}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (sf *StringField) Name() string {
	return sf.name
}

// Value returns the field's value.
func (sf *StringField) Value() string {
	return sf.Get()
}

// DefaultValue returns the field's default value.
func (sf *StringField) DefaultValue() string {
	return sf.defaultValue
}

// Description returns the field's description.
func (sf *StringField) Description() string {
	return sf.options.description(sf.DefaultValue())
}

// Get returns the field value or the default value.
func (sf *StringField) Get() string {
	value := os.Getenv(sf.Name())
	if value == "" {
		if sf.options.required {
			requiredError(sf)
		}
		return sf.defaultValue
	}
	if !sf.options.isAllowedValue(value) {
		unallowedError(sf, value, sf.options.allowedValues)
		return sf.defaultValue
	}
	return value
}
