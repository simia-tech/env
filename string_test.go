package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/simia-tech/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	var (
		optional = env.String("OPTIONAL_FIELD", "abc")
		required = env.String("REQUIRED_FIELD", "abc", env.Required())
		allowed  = env.String("ALLOWED_FIELD", "abc", env.AllowedValues("abc", "def"))
		lastErr  error
	)
	env.ErrorHandler = func(err error) {
		lastErr = err
	}

	tcs := []struct {
		name          string
		field         *env.StringField
		value         string
		expectedValue string
		expectedErr   error
	}{
		{"Value", optional, "def", "def", nil},
		{"DefaultValue", optional, "", "abc", nil},
		{"RequiredAndSet", required, "def", "def", nil},
		{"RequiredNotSet", required, "", "abc", fmt.Errorf("required field REQUIRED_FIELD is not set. using default value 'abc'")},
		{"AllowedValue", allowed, "def", "def", nil},
		{"UnallowedValue", allowed, "ghi", "abc", fmt.Errorf("field ALLOWED_FIELD does not allow value 'ghi' (allowed values are 'abc' and 'def'). using default value 'abc'")},
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
