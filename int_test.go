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

func TestInt(t *testing.T) {
	var (
		optional = env.Int("OPTIONAL_FIELD", 1)
		required = env.Int("REQUIRED_FIELD", 1, env.Required())
		allowed  = env.Int("ALLOWED_FIELD", 1, env.AllowedValues("1", "2", "3"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.IntField
		value         string
		expectedValue int
		expectedErr   error
	}{
		{"Value", optional, "2", 2, nil},
		{"DefaultValue", optional, "", 1, nil},
		{"ParseError", optional, "abc", 1, fmt.Errorf("int field OPTIONAL_FIELD could not be parsed - using default value '1'")},
		{"RequiredAndSet", required, "2", 2, nil},
		{"RequiredNotSet", required, "", 1, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value '1'")},
		{"AllowedValue", allowed, "2", 2, nil},
		{"UnallowedValue", allowed, "4", 1, fmt.Errorf("field ALLOWED_FIELD does not allow value '4' (allowed values are '1', '2' and '3') - using default value '1'")},
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
