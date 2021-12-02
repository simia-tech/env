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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/v2"
)

func TestDuration(t *testing.T) {
	var (
		optional = env.Duration("OPTIONAL_FIELD", time.Minute)
		required = env.Duration("REQUIRED_FIELD", time.Minute, env.Required())
		allowed  = env.Duration("ALLOWED_FIELD", time.Minute, env.AllowedValues("1s", "1m", "1h"))
	)

	testFn := func(field *env.DurationField, value string, expectValue time.Duration, expectErr error) func(*testing.T) {
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

	t.Run("Value", testFn(optional, "1s", time.Second, nil))
	t.Run("DefaultValue", testFn(optional, "", time.Minute, nil))
	t.Run("RequiredAndSet", testFn(required, "1s", time.Second, nil))
	t.Run("RequiredNotSet", testFn(required, "", time.Minute, env.ErrRequiredValueIsMissing))
	t.Run("AllowedValue", testFn(allowed, "1s", time.Second, nil))
	t.Run("UnallowedValue", testFn(allowed, "1ms", time.Minute, env.ErrValueIsNotAllowed))
}
