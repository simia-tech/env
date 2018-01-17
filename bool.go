package env

import (
	"os"
)

const (
	trueValue  = "true"
	falseValue = "false"
)

// BoolField implements a duration field.
type BoolField struct {
	name         string
	defaultValue bool
	options      *options
}

// Bool registers a field of the provided name.
func Bool(name string, defaultValue bool, opts ...Option) *BoolField {
	field := &BoolField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Bool field."}, append(opts, AllowedValues("0", "1", falseValue, trueValue, "no", "yes"))),
	}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (i *BoolField) Name() string {
	return i.name
}

// Value returns the field's value.
func (i *BoolField) Value() string {
	if i.Get() {
		return trueValue
	}
	return falseValue
}

// DefaultValue returns the field's default value.
func (i *BoolField) DefaultValue() string {
	if i.defaultValue {
		return trueValue
	}
	return falseValue
}

// Description returns the field's description.
func (i *BoolField) Description() string {
	return i.options.description(i.DefaultValue())
}

// Get returns the field value or the default value.
func (i *BoolField) Get() bool {
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
	switch text {
	case "1", trueValue, "yes":
		return true
	default:
		return false
	}
}
