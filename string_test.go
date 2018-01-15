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
		optionalString = env.String("OPTIONAL_FIELD", "abc")
		requiredString = env.String("REQUIRED_FIELD", "abc", env.Required())
		lastErr        error
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
		{"Value", optionalString, "def", "def", nil},
		{"DefaultValue", optionalString, "", "abc", nil},
		{"RequiredAndSet", requiredString, "def", "def", nil},
		{"RequiredNotSet", requiredString, "", "abc", fmt.Errorf("required field REQUIRED_FIELD is not set. using default value 'abc'")},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, os.Setenv(tc.field.Name(), tc.value))

			assert.Equal(t, tc.expectedValue, tc.field.Get())

			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, lastErr)
			}
		})
	}
}
