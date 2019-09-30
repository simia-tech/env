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

func TestBool(t *testing.T) {
	var (
		optional = env.Bool("OPTIONAL_FIELD", false)
		required = env.Bool("REQUIRED_FIELD", false, env.Required())
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	testFn := func(field *env.BoolField, value string, expectValue bool, expectErr error) func(*testing.T) {
		return func(t *testing.T) {
			require.NoError(t, os.Setenv(field.Name(), value))

			assert.Equal(t, expectValue, field.Get())

			if expectErr != nil {
				assert.Equal(t, expectErr, lastErr)
				lastErr = nil
			}
		}
	}

	t.Run("Value", testFn(optional, "1", true, nil))
	t.Run("DefaultValue", testFn(optional, "", false, nil))
	t.Run("RequiredAndSet", testFn(required, "yes", true, nil))
	t.Run("RequiredNotSet", testFn(required, "", false, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value 'false'")))
	t.Run("UnallowedValue", testFn(optional, "okaydokay", false, fmt.Errorf("field OPTIONAL_FIELD does not allow value 'okaydokay' (allowed values are '0', '1', 'false', 'true', 'no' and 'yes') - using default value 'false'")))
}
