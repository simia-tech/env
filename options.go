package env

// Option defines an Option that can modify the options struct.
type Option func(*options)

type options struct {
	required      bool
	allowedValues []string
	description   string
}

func newOptions(opts []Option) *options {
	o := &options{}
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
		o.description = text
	}
}

func isAllowedValue(options *options, value string) bool {
	if options.allowedValues == nil {
		return true
	}
	for _, allowedValue := range options.allowedValues {
		if value == allowedValue {
			return true
		}
	}
	return false
}
