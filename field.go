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
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// Field implements an environment configuration field.
type Field interface {
	Name() string
	Value() string
	DefaultValue() string
	Description() string
}

var fields = map[string]Field{}
var nameRegexp = regexp.MustCompile("^[A-Z0-9_]+$")

// RegisterField adds the provided `Field` to the global field-register.
func RegisterField(field Field) Field {
	name := field.Name()
	if !nameRegexp.MatchString(name) {
		panic(fmt.Sprintf("field name [%s] must only contain capital letters, numbers or underscores", name))
	}
	if f, ok := fields[name]; ok {
		return f
	}
	fields[name] = field
	return field
}

// Fields returns a slice of strings with all registered fields.
func Fields() []string {
	names := []string{}
	for name := range fields {
		names = append(names, name)
	}
	return names
}

// Clear clears the field register.
func Clear() {
	fields = map[string]Field{}
}

type field struct {
	label    string
	name     string
	location string
	options  *options
}

func newField(label, name string, opts []Option) field {
	_, filename, line, _ := runtime.Caller(2)
	return field{
		label:    label,
		name:     name,
		location: fmt.Sprintf("%s:%d", filename, line),
		options:  newOptions(opts),
	}
}

func (f *field) Name() string {
	return f.name
}

func (f *field) value() (string, error) {
	value := os.Getenv(f.name)
	if f.options.required && value == "" {
		return "", fmt.Errorf("field %s: %w", f.name, ErrRequiredValueIsMissing)
	}
	if !f.options.isAllowedValue(value) {
		return "", fmt.Errorf("field %s with value [%s]: %w", f.name, value, ErrValueIsNotAllowed)
	}
	return value, nil
}

func (f *field) description(defaultValue string) string {
	if f.options.desc != "" {
		return f.options.desc
	}
	sentences := []string{f.label + " field."}
	if f.options.required {
		sentences = append(sentences, "Required field.")
	}
	if f.options.allowedValues != nil {
		sentences = append(sentences, fmt.Sprintf("Allowed values are %s.", joinStringValues(f.options.allowedValues)))
	}
	sentences = append(sentences, "The default value is '"+defaultValue+"'.")
	sentences = append(sentences, "Defined at "+f.location+".")
	return strings.Join(sentences, " ")
}
