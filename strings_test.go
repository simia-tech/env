package env_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/simia-tech/env"
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
