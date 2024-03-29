// Copyright 2018 Philipp Brüll <pb@simia.tech>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrRequiredValueIsMissing = errors.New("required value is missing")
	ErrValueIsNotAllowed      = errors.New("value is not allowed")
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
