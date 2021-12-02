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
	"time"
)

// DurationField implements a duration field.
type DurationField struct {
	field
	defaultValue time.Duration
}

// Duration registers a field of the provided name.
func Duration(name string, defaultValue time.Duration, opts ...Option) *DurationField {
	field := &DurationField{
		field:        newField("Duration", name, opts),
		defaultValue: defaultValue,
	}
	RegisterField(field)
	return field
}

// Value returns the field's value.
func (f *DurationField) Value() string {
	return f.GetOrDefault().String()
}

// DefaultValue returns the field's default value.
func (f *DurationField) DefaultValue() string {
	return f.defaultValue.String()
}

// Description returns the field's description.
func (f *DurationField) Description() string {
	return f.description(f.DefaultValue())
}

// GetOrDefault returns the field value or the default value.
func (f *DurationField) GetOrDefault() time.Duration {
	value, err := f.Get()
	if err != nil {
		ErrorHandler(err)
		return f.defaultValue
	}
	return value
}

// Get returns the field value or an error.
func (f *DurationField) Get() (time.Duration, error) {
	v, err := f.value()
	if err != nil {
		return f.defaultValue, err
	}
	if v == "" {
		return f.defaultValue, nil
	}
	value, err := time.ParseDuration(v)
	if err != nil {
		return f.defaultValue, fmt.Errorf("field %s parse duration [%s]: %w", f.name, v, err)
	}
	return value, nil
}
