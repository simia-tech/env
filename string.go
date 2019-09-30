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

import "os"

// StringField implements a string field.
type StringField struct {
	*field
	defaultValue string
}

// String registers a field of the provided name.
func String(name, defaultValue string, opts ...Option) *StringField {
	field := &StringField{
		field:        newField("String", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (sf *StringField) Value() string {
	return sf.Get()
}

// DefaultValue returns the field's default value.
func (sf *StringField) DefaultValue() string {
	return sf.defaultValue
}

// Description returns the field's description.
func (sf *StringField) Description() string {
	return sf.description(sf.DefaultValue())
}

// Get returns the field value or the default value.
func (sf *StringField) Get() string {
	value := os.Getenv(sf.Name())
	if value == "" {
		if sf.options.required {
			requiredError(sf)
		}
		return sf.defaultValue
	}
	if !sf.options.isAllowedValue(value) {
		unallowedError(sf, value, sf.options.allowedValues)
		return sf.defaultValue
	}
	return value
}
