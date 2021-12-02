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
	"strings"
)

const separator = ","

// StringsField implements a strings field.
type StringsField struct {
	*field
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
func (sf *StringsField) Value() string {
	return strings.Join(sf.Get(), separator)
}

// DefaultValue returns the field's default value.
func (sf *StringsField) DefaultValue() string {
	return strings.Join(sf.defaultValue, separator)
}

// Description returns the field's description.
func (sf *StringsField) Description() string {
	return sf.description(sf.DefaultValue())
}

// Get returns the field value or the default value.
func (sf *StringsField) Get() []string {
	text := os.Getenv(sf.Name())
	if text == "" {
		if sf.options.required {
			requiredError(sf)
		}
		return sf.defaultValue
	}
	values := strings.Split(text, separator)
	for _, value := range values {
		if !sf.options.isAllowedValue(value) {
			unallowedError(sf, value, sf.options.allowedValues)
			return sf.defaultValue
		}
	}
	return values
}
