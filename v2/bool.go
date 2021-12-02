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

const (
	trueValue  = "true"
	falseValue = "false"
)

// BoolField implements a duration field.
type BoolField struct {
	field
	defaultValue bool
}

// Bool registers a field of the provided name.
func Bool(name string, defaultValue bool, opts ...Option) *BoolField {
	field := &BoolField{
		field:        newField("Boolean", name, append(opts, AllowedValues("0", "1", falseValue, trueValue, "no", "yes", ""))),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *BoolField) Value() string {
	if f.GetOrDefault() {
		return trueValue
	}
	return falseValue
}

// DefaultValue returns the field's default value.
func (f *BoolField) DefaultValue() string {
	if f.defaultValue {
		return trueValue
	}
	return falseValue
}

// Description returns the field's description.
func (f *BoolField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default in case of an error.
func (f *BoolField) GetOrDefault() bool {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error.
func (f *BoolField) Get() (bool, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	switch v {
	case "1", trueValue, "yes":
		return true, nil
	default:
		return false, nil
	}
}
