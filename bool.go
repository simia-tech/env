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
)

const (
	trueValue  = "true"
	falseValue = "false"
)

// BoolField implements a duration field.
type BoolField struct {
	name         string
	defaultValue bool
	options      *options
}

// Bool registers a field of the provided name.
func Bool(name string, defaultValue bool, opts ...Option) *BoolField {
	field := &BoolField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Bool field."}, append(opts, AllowedValues("0", "1", falseValue, trueValue, "no", "yes"))),
	}
	RegisterField(field)
	return field
}

// Name returns the field name.
func (i *BoolField) Name() string {
	return i.name
}

// Value returns the field's value.
func (i *BoolField) Value() string {
	if i.Get() {
		return trueValue
	}
	return falseValue
}

// DefaultValue returns the field's default value.
func (i *BoolField) DefaultValue() string {
	if i.defaultValue {
		return trueValue
	}
	return falseValue
}

// Description returns the field's description.
func (i *BoolField) Description() string {
	return i.options.description(i.DefaultValue())
}

// Get returns the field value or the default value.
func (i *BoolField) Get() bool {
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
	switch text {
	case "1", trueValue, "yes":
		return true
	default:
		return false
	}
}
