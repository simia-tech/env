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

func TestInts(t *testing.T) {
	var (
		optional = env.Ints("OPTIONAL_FIELD", []int{123})
		required = env.Ints("REQUIRED_FIELD", []int{123}, env.Required())
		allowed  = env.Ints("ALLOWED_FIELD", []int{123}, env.AllowedValues("123", "456"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.IntsField
		value         string
		expectedValue []int
		expectedErr   error
	}{
		{"Value", optional, "456", []int{456}, nil},
		{"DefaultValue", optional, "", []int{123}, nil},
		{"RequiredAndSet", required, "456", []int{456}, nil},
		{"RequiredNotSet", required, "", []int{123}, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value '123'")},
		{"AllowedValue", allowed, "456", []int{456}, nil},
		{"UnallowedValue", allowed, "789", []int{123}, fmt.Errorf("field ALLOWED_FIELD does not allow value '789' (allowed values are '123' and '456') - using default value '123'")},
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
