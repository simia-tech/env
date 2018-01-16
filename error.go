package env

import (
	"fmt"
	"log"
	"os"
)

// ErrorHandler defines a handler for error messages. By default, LogErrorHandler is set.
var ErrorHandler = StderrErrorHandler

// Implementations of different error handlers.
var (
	NullErrorHandler   = func(error) {}
	StdoutErrorHandler = func(err error) {
		fmt.Fprintln(os.Stdout, err)
	}
	StderrErrorHandler = func(err error) {
		fmt.Fprintln(os.Stderr, err)
	}
	LogErrorHandler = func(err error) {
		log.Println(err)
	}
)

func requiredError(field Field) {
	ErrorHandler(fmt.Errorf("required field %s is not set - using default value '%s'", field.Name(), field.DefaultValue()))
}

func parseError(field Field, kind, value string) {
	ErrorHandler(fmt.Errorf("%s field %s could not be parsed - using default value '%s'", kind, field.Name(), field.DefaultValue()))
}

func unallowedError(field Field, value string, allowedValues []string) {
	if len(allowedValues) == 1 {
		ErrorHandler(
			fmt.Errorf("field %s does not allow value '%s' (only value '%s' is allowed) - using default value '%s'",
				field.Name(), value, allowedValues[0], field.DefaultValue()))
	} else {
		ErrorHandler(
			fmt.Errorf("field %s does not allow value '%s' (allowed values are %s) - using default value '%s'",
				field.Name(), value, joinStringValues(allowedValues), field.DefaultValue()))
	}
}

func joinStringValues(values []string) string {
	return joinStrings(values, ", ", " and ")
}

func joinStrings(values []string, sepRune, sepWord string) string {
	text := ""
	for index := 0; index < len(values)-1; index++ {
		if index > 0 {
			text += sepRune
		}
		text += "'" + values[index] + "'"
	}
	text += sepWord + "'" + values[len(values)-1] + "'"
	return text
}
