package env

import (
	"fmt"
	"log"
)

// ErrorHandler defines a handler for error messages. By default, LogErrorHandler is set.
var ErrorHandler = LogErrorHandler

// LogErrorHandler defines a error handler that prints all errors to the log.
var LogErrorHandler = func(err error) {
	log.Println(err)
}

func requiredError(name, defaultValue string) {
	ErrorHandler(fmt.Errorf("required field %s is not set. using default value '%s'", name, defaultValue))
}

func parseError(kind, name, value, defaultValue string) {
	ErrorHandler(fmt.Errorf("%s field %s could not be parsed. using default value '%s'", kind, name, defaultValue))
}
