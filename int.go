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
	"strconv"
)

// IntField implements a duration field.
type IntField struct {
	field
	defaultValue int
}

// Int registers a field of the provided name.
func Int(name string, defaultValue int, opts ...Option) *IntField {
	field := &IntField{
		field:        newField("Int", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *IntField) Value() string {
	return strconv.Itoa(f.GetOrDefault())
}

// DefaultValue returns the field's default value.
func (f *IntField) DefaultValue() string {
	return strconv.Itoa(f.defaultValue)
}

// Description returns the field's description.
func (f *IntField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *IntField) GetOrDefault() int {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error.
func (f *IntField) Get() (int, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return f.defaultValue, fmt.Errorf("field %s parse int [%s]: %w", f.name, v, err)
	}
	return value, nil
}
