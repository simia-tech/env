package env

import (
	"fmt"
	"strings"
)

// Option defines an Option that can modify the options struct.
type Option func(*options)

type options struct {
	required      bool
	allowedValues []string
	desc          string
	descSentences []string
}

func newOptions(descSentences []string, opts []Option) *options {
	o := &options{descSentences: descSentences}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Required returns an Option that makes the environment field required.
func Required() Option {
	return func(o *options) {
		o.required = true
	}
}

// AllowedValues returns an Option that defines all allowed values for the environment field.
func AllowedValues(values ...string) Option {
	return func(o *options) {
		o.allowedValues = values
	}
}

// Description returns an Option that sets the description of the environment field.
func Description(text string) Option {
	return func(o *options) {
		o.desc = text
	}
}

func (o *options) isAllowedValue(value string) bool {
	if o == nil || o.allowedValues == nil {
		return true
	}
	for _, allowedValue := range o.allowedValues {
		if value == allowedValue {
			return true
		}
	}
	return false
}

func (o *options) description(defaultValue string) string {
	if o == nil {
		return ""
	}
	if o.desc != "" {
		return o.desc
	}
	sentences := o.descSentences
	if o.required {
		sentences = append(sentences, "Required field.")
	}
	if o.allowedValues != nil {
		sentences = append(sentences, fmt.Sprintf("Allowed values are %s.", joinStringValues(o.allowedValues)))
	}
	sentences = append(sentences, "The default value is '"+defaultValue+"'.")
	return strings.Join(sentences, " ")
}
