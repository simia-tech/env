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
	"time"
)

// DurationField implements a duration field.
type DurationField struct {
	name         string
	defaultValue time.Duration
	options      *options
}

// Duration registers a field of the provided name.
func Duration(name string, defaultValue time.Duration, opts ...Option) *DurationField {
	field := &DurationField{
		name:         name,
		defaultValue: defaultValue,
		options:      newOptions([]string{"Duration field."}, opts),
	}
	fields = append(fields, field)
	return field
}

// Name returns the field name.
func (df *DurationField) Name() string {
	return df.name
}

// Value returns the field's value.
func (df *DurationField) Value() string {
	return df.Get().String()
}

// DefaultValue returns the field's default value.
func (df *DurationField) DefaultValue() string {
	return df.defaultValue.String()
}

// Description returns the field's description.
func (df *DurationField) Description() string {
	return df.options.description(df.DefaultValue())
}

// Get returns the field value or the default value.
func (df *DurationField) Get() time.Duration {
	text := os.Getenv(df.Name())
	if text == "" {
		if df.options.required {
			requiredError(df)
		}
		return df.defaultValue
	}
	if !df.options.isAllowedValue(text) {
		unallowedError(df, text, df.options.allowedValues)
		return df.defaultValue
	}
	value, err := time.ParseDuration(text)
	if err != nil {
		parseError(df, "duration", text)
		return df.defaultValue
	}
	return value
}
