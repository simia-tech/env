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

	"github.com/simia-tech/env/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrings(t *testing.T) {
	var (
		optional = env.Strings("OPTIONAL_FIELD", []string{"abc"})
		required = env.Strings("REQUIRED_FIELD", []string{"abc"}, env.Required())
		allowed  = env.Strings("ALLOWED_FIELD", []string{"abc"}, env.AllowedValues("abc", "def"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.StringsField
		value         string
		expectedValue []string
		expectedErr   error
	}{
		{"Value", optional, "def", []string{"def"}, nil},
		{"DefaultValue", optional, "", []string{"abc"}, nil},
		{"RequiredAndSet", required, "def", []string{"def"}, nil},
		{"RequiredNotSet", required, "", []string{"abc"}, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value 'abc'")},
		{"AllowedValue", allowed, "def", []string{"def"}, nil},
		{"UnallowedValue", allowed, "ghi", []string{"abc"}, fmt.Errorf("field ALLOWED_FIELD does not allow value 'ghi' (allowed values are 'abc' and 'def') - using default value 'abc'")},
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
