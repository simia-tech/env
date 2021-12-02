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

	"github.com/simia-tech/env/v2/internal/parser"
)

// StringMapField implements a map of strings.
type StringMapField struct {
	field
	defaultValue map[string]string
}

// StringMap registers a field of the provided name.
func StringMap(name string, defaultValue map[string]string, opts ...Option) *StringMapField {
	field := &StringMapField{
		field:        newField("Strings", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *StringMapField) Value() string {
	return parser.FormatStringMap(f.GetOrDefault())
}

// DefaultValue returns the field's default value.
func (f *StringMapField) DefaultValue() string {
	return parser.FormatStringMap(f.defaultValue)
}

// Description returns the field's description.
func (f *StringMapField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *StringMapField) GetOrDefault() map[string]string {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error
func (f *StringMapField) Get() (map[string]string, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}

	values, err := parser.ParseStringMap(v)
	if err != nil {
		return f.defaultValue, fmt.Errorf("parse string map: %w", err)
	}
	return values, nil
}
