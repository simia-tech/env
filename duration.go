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

// Get returns the field value or the default value.
func (df *DurationField) Get() time.Duration {
	text := os.Getenv(df.name)
	if text == "" {
		if df.options.required {
			requiredError(df.name, df.defaultValue.String())
		}
		return df.defaultValue
	}
	value, err := time.ParseDuration(text)
	if err != nil {
		parseError("duration", df.name, text, df.defaultValue.String())
		return df.defaultValue
	}
	return value
}
