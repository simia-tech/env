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

package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env"
)

func TestStringMap(t *testing.T) {
	var (
		optional = env.StringMap("OPTIONAL_FIELD", map[string]string{"abc": "123"})
		required = env.StringMap("REQUIRED_FIELD", map[string]string{"abc": "123"}, env.Required())
		allowed  = env.StringMap("ALLOWED_FIELD", map[string]string{"abc": "def"}, env.AllowedValues("abc", "def"))
	)

	testFn := func(field *env.StringMapField, value string, expectValue map[string]string, expectErr error) func(*testing.T) {
		return func(t *testing.T) {
			require.NoError(t, os.Setenv(field.Name(), value))

			value, err := field.Get()
			if expectErr == nil {
				require.NoError(t, err)
				assert.Equal(t, expectValue, value)
			} else {
				assert.ErrorIs(t, err, expectErr)
			}
		}
	}

	t.Run("Value", testFn(optional, "def", map[string]string{"def": ""}, nil))
	t.Run("DefaultValue", testFn(optional, "", map[string]string{"abc": "123"}, nil))
	t.Run("RequiredAndSet", testFn(required, "def", map[string]string{"def": ""}, nil))
	t.Run("RequiredNotSet", testFn(required, "", map[string]string{"abc": "123"}, env.ErrRequiredValueIsMissing))
	t.Run("AllowedValue", testFn(allowed, "def", map[string]string{"def": ""}, nil))
	t.Run("UnallowedValue", testFn(allowed, "ghi", map[string]string{"abc": "def"}, env.ErrValueIsNotAllowed))
}
