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
		{"ParseError", optional, "abc", time.Minute, fmt.Errorf("duration field OPTIONAL_FIELD could not be parsed. using default value '1m0s'")},
		{"RequiredAndSet", required, "1s", time.Second, nil},
		{"RequiredNotSet", required, "", time.Minute, fmt.Errorf("required field REQUIRED_FIELD is not set. using default value '1m0s'")},
		{"AllowedValue", allowed, "1s", time.Second, nil},
		{"UnallowedValue", allowed, "1ms", time.Minute, fmt.Errorf("field ALLOWED_FIELD does not allow value '1ms' (allowed values are '1s', '1m' and '1h'). using default value '1m0s'")},
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
