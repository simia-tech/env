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
	"fmt"
)

// BytesField implements a string field.
type BytesField struct {
	field
	defaultValue []byte
}

// Bytes registers a field of the provided name.
func Bytes(name string, defaultValue []byte, opts ...Option) *BytesField {
	field := &BytesField{
		field:        newField("Bytes", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *BytesField) Value() string {
	return hex.EncodeToString(f.GetOrDefault())
}

// DefaultValue returns the field's default value.
func (f *BytesField) DefaultValue() string {
	return hex.EncodeToString(f.defaultValue)
}

// Description returns the field's description.
func (f *BytesField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *BytesField) GetOrDefault() []byte {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error.
func (f *BytesField) Get() ([]byte, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}
	value, err := hex.DecodeString(v)
	if err != nil {
		return f.defaultValue, fmt.Errorf("field %s hex decode [%s]: %w", f.name, v, err)
	}
	return value, nil
}
