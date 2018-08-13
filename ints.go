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
	"os"
	"strconv"
	"strings"
)

// IntsField implements a ints field.
type IntsField struct {
	name         string
	defaultValue []int
	options      *options
}

// Ints registers a field of the provided name.
func Ints(name string, defaultValue []int, opts ...Option) *IntsField {
	field := &IntsField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Strings field."}, opts),
	}
	RegisterField(field)
	return field
}

// Name returns the field name.
func (isf *IntsField) Name() string {
	return isf.name
}

// Value returns the field's value.
func (isf *IntsField) Value() string {
	return strings.Join(stringsValue(isf.Get()), separator)
}

// DefaultValue returns the field's default value.
func (isf *IntsField) DefaultValue() string {
	return strings.Join(stringsValue(isf.defaultValue), separator)
}

// Description returns the field's description.
func (isf *IntsField) Description() string {
	return isf.options.description(isf.DefaultValue())
}

// Get returns the field value or the default value.
func (isf *IntsField) Get() []int {
	text := os.Getenv(isf.Name())
	if text == "" {
		if isf.options.required {
			requiredError(isf)
		}
		return isf.defaultValue
	}
	values := []int{}
	for _, text := range strings.Split(text, separator) {
		if !isf.options.isAllowedValue(text) {
			unallowedError(isf, text, isf.options.allowedValues)
			return isf.defaultValue
		}
		value, err := strconv.Atoi(text)
		if err != nil {
			parseError(isf, "ints", text)
			return isf.defaultValue
		}
		values = append(values, value)
	}
	return values
}

func stringsValue(values []int) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = strconv.Itoa(value)
	}
	return result
}
