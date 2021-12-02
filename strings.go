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
	"strings"

	"github.com/simia-tech/env/v2/internal/parser"
)

const separator = ","

// StringsField implements a strings field.
type StringsField struct {
	field
	defaultValue []string
}

// Strings registers a field of the provided name.
func Strings(name string, defaultValue []string, opts ...Option) *StringsField {
	field := &StringsField{
		field:        newField("Strings", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *StringsField) Value() string {
	return strings.Join(f.GetOrDefault(), separator)
}

// DefaultValue returns the field's default value.
func (f *StringsField) DefaultValue() string {
	return strings.Join(f.defaultValue, separator)
}

// Description returns the field's description.
func (f *StringsField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *StringsField) GetOrDefault() []string {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error
func (f *StringsField) Get() ([]string, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}

	values, err := parser.ParseStrings(v)
	if err != nil {
		return f.defaultValue, err
	}
	return values, nil
}
