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

func unallowedError(name, value string, allowedValues []string, defaultValue string) {
	switch length := len(allowedValues); length {
	case 1:
		ErrorHandler(fmt.Errorf("field %s does not allow value '%s' (only value '%s' is allowed). using default value '%s'", name, value, allowedValues[0], defaultValue))
	default:
		text := ""
		for index := 0; index < length-1; index++ {
			if index > 0 {
				text += ", "
			}
			text += "'" + allowedValues[index] + "'"
		}
		text += " and '" + allowedValues[length-1] + "'"
		ErrorHandler(fmt.Errorf("field %s does not allow value '%s' (allowed values are %s). using default value '%s'", name, value, text, defaultValue))
	}
}
