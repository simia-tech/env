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
	"regexp"
)

// Field implements an environment configuration field.
type Field interface {
	Name() string
	Value() string
	DefaultValue() string
	Description() string
}

var fields = map[string]Field{}
var nameRegexp = regexp.MustCompile("")

// RegisterField adds the provided `Field` to the global field-register.
func RegisterField(field Field) Field {
	name := field.Name()
	if !nameRegexp.MatchString(name) {
		panic(fmt.Sprintf("field name [%s] must only contain capital letters, numbers or underscores", name))
	}
	if f, ok := fields[name]; ok {
		return f
	}
	fields[name] = field
	return field
}

// Fields returns a slice of strings with all registered fields.
func Fields() []string {
	names := []string{}
	for name := range fields {
		names = append(names, name)
	}
	return names
}

// Clear clears the field register.
func Clear() {
	fields = map[string]Field{}
}
