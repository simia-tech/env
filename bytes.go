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
	"encoding/hex"
	"os"
)

// BytesField implements a string field.
type BytesField struct {
	name         string
	defaultValue []byte
	options      *options
}

// Bytes registers a field of the provided name.
func Bytes(name string, defaultValue []byte, opts ...Option) *BytesField {
	field := &BytesField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Bytes field."}, opts),
	}
	RegisterField(field)
	return field
}

// Name returns the field name.
func (bf *BytesField) Name() string {
	return bf.name
}

// Value returns the field's value.
func (bf *BytesField) Value() string {
	return hex.EncodeToString(bf.Get())
}

// DefaultValue returns the field's default value.
func (bf *BytesField) DefaultValue() string {
	return hex.EncodeToString(bf.defaultValue)
}

// Description returns the field's description.
func (bf *BytesField) Description() string {
	return bf.options.description(bf.DefaultValue())
}

// Get returns the field value or the default value.
func (bf *BytesField) Get() []byte {
	text := os.Getenv(bf.Name())
	if text == "" {
		if bf.options.required {
			requiredError(bf)
		}
		return bf.defaultValue
	}
	if !bf.options.isAllowedValue(text) {
		unallowedError(bf, text, bf.options.allowedValues)
		return bf.defaultValue
	}
	value, err := hex.DecodeString(text)
	if err != nil {
		parseError(bf, "bytes", text)
		return bf.defaultValue
	}
	return value
}
