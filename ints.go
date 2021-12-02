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
	"strings"
)

// IntsField implements a ints field.
type IntsField struct {
	field
	defaultValue []int
}

// Ints registers a field of the provided name.
func Ints(name string, defaultValue []int, opts ...Option) *IntsField {
	field := &IntsField{
		field:        newField("Ints", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *IntsField) Value() string {
	return strings.Join(stringsValue(f.GetOrDefault()), separator)
}

// DefaultValue returns the field's default value.
func (f *IntsField) DefaultValue() string {
	return strings.Join(stringsValue(f.defaultValue), separator)
}

// Description returns the field's description.
func (f *IntsField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *IntsField) GetOrDefault() []int {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error.
func (f *IntsField) Get() ([]int, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}

	parts := strings.Split(v, separator)
	values := make([]int, len(parts))
	for index, v := range parts {
		if !f.options.isAllowedValue(v) {
			return f.defaultValue, fmt.Errorf("field %s.%d with value [%s]: %w", f.name, index, v, ErrValueIsNotAllowed)
		}
		value, err := strconv.Atoi(v)
		if err != nil {
			return f.defaultValue, fmt.Errorf("field %s.%d parse int [%s]: %w", f.name, index, v, err)
		}
		values[index] = value
	}
	return values, nil
}

func stringsValue(values []int) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = strconv.Itoa(value)
	}
	return result
}
