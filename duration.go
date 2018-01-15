package env

import (
	"os"
	"time"
)

// DurationField implements a duration field.
type DurationField struct {
	name         string
	defaultValue time.Duration
	options      *options
}

// Duration registers a field of the provided name.
func Duration(name string, defaultValue time.Duration, opts ...Option) *DurationField {
	field := &DurationField{name: name, defaultValue: defaultValue, options: newOptions(opts)}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (df *DurationField) Name() string {
	return df.name
}

// DefaultValue returns the field's default value.
func (df *DurationField) DefaultValue() string {
	return df.defaultValue.String()
}

// Description returns the field's description.
func (df *DurationField) Description() string {
	return df.options.description()
}

// Get returns the field value or the default value.
func (df *DurationField) Get() time.Duration {
	text := os.Getenv(df.name)
	if text == "" {
		if df.options.required {
			requiredError(df)
		}
		return df.defaultValue
	}
	if !df.options.isAllowedValue(text) {
		unallowedError(df, text, df.options.allowedValues)
		return df.defaultValue
	}
	value, err := time.ParseDuration(text)
	if err != nil {
		parseError(df, "duration", text)
		return df.defaultValue
	}
	return value
}
