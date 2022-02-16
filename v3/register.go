package env

var fields = []generalField{}

type generalField interface {
	Name() string
	Description() string
	GetRaw() (string, error)
	GetRawOrDefault() string
}

func ClearRegister() {
	fields = []generalField{}
}

func registerField(field generalField) {
	fields = append(fields, field)
}
