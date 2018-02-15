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
	"fmt"
	"os"
	"testing"

	"github.com/simia-tech/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBytes(t *testing.T) {
	var (
		optional = env.Bytes("OPTIONAL_FIELD", []byte{0xab})
		required = env.Bytes("REQUIRED_FIELD", []byte{0xab}, env.Required())
		allowed  = env.Bytes("ALLOWED_FIELD", []byte{0xab}, env.AllowedValues("ab", "cd"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.BytesField
		value         string
		expectedValue []byte
		expectedErr   error
	}{
		{"Value", optional, "cd", []byte{0xcd}, nil},
		{"DefaultValue", optional, "", []byte{0xab}, nil},
		{"RequiredAndSet", required, "cd", []byte{0xcd}, nil},
		{"RequiredNotSet", required, "", []byte{0xab}, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value 'ab'")},
		{"AllowedValue", allowed, "cd", []byte{0xcd}, nil},
		{"UnallowedValue", allowed, "ef", []byte{0xab}, fmt.Errorf("field ALLOWED_FIELD does not allow value 'ef' (allowed values are 'ab' and 'cd') - using default value 'ab'")},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, os.Setenv(tc.field.Name(), tc.value))

			assert.Equal(t, tc.expectedValue, tc.field.Get())

			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, lastErr)
				lastErr = nil
			}
		})
	}
}
