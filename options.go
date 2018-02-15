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
	"strings"
)

// Option defines an Option that can modify the options struct.
type Option func(*options)

type options struct {
	required      bool
	allowedValues []string
	desc          string
	descSentences []string
}

func newOptions(descSentences []string, opts []Option) *options {
	o := &options{descSentences: descSentences}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Required returns an Option that makes the environment field required.
func Required() Option {
	return func(o *options) {
		o.required = true
	}
}

// AllowedValues returns an Option that defines all allowed values for the environment field.
func AllowedValues(values ...string) Option {
	return func(o *options) {
		o.allowedValues = values
	}
}

// Description returns an Option that sets the description of the environment field.
func Description(text string) Option {
	return func(o *options) {
		o.desc = text
	}
}

func (o *options) isAllowedValue(value string) bool {
	if o == nil || o.allowedValues == nil {
		return true
	}
	for _, allowedValue := range o.allowedValues {
		if value == allowedValue {
			return true
		}
	}
	return false
}

func (o *options) description(defaultValue string) string {
	if o == nil {
		return ""
	}
	if o.desc != "" {
		return o.desc
	}
	sentences := o.descSentences
	if o.required {
		sentences = append(sentences, "Required field.")
	}
	if o.allowedValues != nil {
		sentences = append(sentences, fmt.Sprintf("Allowed values are %s.", joinStringValues(o.allowedValues)))
	}
	sentences = append(sentences, "The default value is '"+defaultValue+"'.")
	return strings.Join(sentences, " ")
}
