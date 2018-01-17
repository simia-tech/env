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

	tcs := []struct {
		name          string
		field         *env.BoolField
		value         string
		expectedValue bool
		expectedErr   error
	}{
		{"Value", optional, "1", true, nil},
		{"DefaultValue", optional, "", false, nil},
		{"RequiredAndSet", required, "yes", true, nil},
		{"RequiredNotSet", required, "", false, fmt.Errorf("required field REQUIRED_FIELD is not set - using default value 'false'")},
		{"UnallowedValue", optional, "okaydokay", false, fmt.Errorf("field OPTIONAL_FIELD does not allow value 'okaydokay' (allowed values are '0', '1', 'false', 'true', 'no' and 'yes') - using default value 'false'")},
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
