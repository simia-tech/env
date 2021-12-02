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

func TestBytes(t *testing.T) {
	var (
		optional = env.Bytes("OPTIONAL_FIELD", []byte{0xab})
		required = env.Bytes("REQUIRED_FIELD", []byte{0xab}, env.Required())
		allowed  = env.Bytes("ALLOWED_FIELD", []byte{0xab}, env.AllowedValues("ab", "cd"))
	)

	testFn := func(field *env.BytesField, value string, expectValue []byte, expectErr error) func(*testing.T) {
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

	t.Run("Value", testFn(optional, "cd", []byte{0xcd}, nil))
	t.Run("DefaultValue", testFn(optional, "", []byte{0xab}, nil))
	t.Run("RequiredAndSet", testFn(required, "cd", []byte{0xcd}, nil))
	t.Run("RequiredNotSet", testFn(required, "", []byte{0xab}, env.ErrRequiredValueIsMissing))
	t.Run("AllowedValue", testFn(allowed, "cd", []byte{0xcd}, nil))
	t.Run("UnallowedValue", testFn(allowed, "ef", []byte{0xab}, env.ErrValueIsNotAllowed))
}
