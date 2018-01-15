package env

// Option defines an Option that can modify the options struct.
type Option func(*options)

type options struct {
	required    bool
	description string
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Required returns an Option that makes the environment field required.
func Required() func(*options) {
	return func(o *options) {
		o.required = true
	}
}

// Description returns an Option that sets the description of the environment field.
func Description(text string) func(*options) {
	return func(o *options) {
		o.description = text
	}
}
