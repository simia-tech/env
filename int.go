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
)

// IntField implements a duration field.
type IntField struct {
	*field
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
func (i *IntField) Value() string {
	return strconv.Itoa(i.Get())
}

// DefaultValue returns the field's default value.
func (i *IntField) DefaultValue() string {
	return strconv.Itoa(i.defaultValue)
}

// Description returns the field's description.
func (i *IntField) Description() string {
	return i.description(i.DefaultValue())
}

// Get returns the field value or the default value.
func (i *IntField) Get() int {
	text := os.Getenv(i.Name())
	if text == "" {
		if i.options.required {
			requiredError(i)
		}
		return i.defaultValue
	}
	if !i.options.isAllowedValue(text) {
		unallowedError(i, text, i.options.allowedValues)
		return i.defaultValue
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		parseError(i, "int", text)
		return i.defaultValue
	}
	return value
}
