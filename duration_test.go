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
	"time"

	"github.com/simia-tech/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration(t *testing.T) {
	var (
		optional = env.Duration("OPTIONAL_FIELD", time.Minute)
		required = env.Duration("REQUIRED_FIELD", time.Minute, env.Required())
		allowed  = env.Duration("ALLOWED_FIELD", time.Minute, env.AllowedValues("1s", "1m", "1h"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.DurationField
		value         string
		expectedValue time.Duration
		expectedErr   error
	}{
		{"Value", optional, "1s", time.Second, nil},
		{"DefaultValue", optional, "", time.Minute, nil},
		{"ParseError", optional, "abc", time.Minute, fmt.Errorf("duration field OPTIONAL_FIELD could not be parsed - using default value '1m0s'")},
		{"RequiredAndSet", required, "1s", time.Second, nil},
		{"RequiredNotSet", required, "", time.Minute, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value '1m0s'")},
		{"AllowedValue", allowed, "1s", time.Second, nil},
		{"UnallowedValue", allowed, "1ms", time.Minute, fmt.Errorf("field ALLOWED_FIELD does not allow value '1ms' (allowed values are '1s', '1m' and '1h') - using default value '1m0s'")},
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
