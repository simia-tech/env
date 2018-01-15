package env

// Field implements an environment configuration field.
type Field interface {
	Name() string
}

var fields = []Field{}

// Fields returns a slice of strings with all registered fields.
func Fields() []string {
	names := []string{}
	for _, field := range fields {
		names = append(names, field.Name())
	}
	return names
}
