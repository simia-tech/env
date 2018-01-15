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
	field := &StringField{name: name, defaultValue: defaultValue, options: newOptions(opts)}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (sf *StringField) Name() string {
	return sf.name
}

// Get returns the field value or the default value.
func (sf *StringField) Get() string {
	value := os.Getenv(sf.name)
	if value == "" {
		if sf.options.required {
			requiredError(sf.name, sf.defaultValue)
		}
		return sf.defaultValue
	}
	if !isAllowedValue(sf.options, value) {
		unallowedError(sf.name, value, sf.options.allowedValues, sf.defaultValue)
		return sf.defaultValue
	}
	return value
}
