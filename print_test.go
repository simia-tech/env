package env_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env"
)

func TestPrint(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Required(), env.AllowedValues("one", "two", "three"))

	tcs := []struct {
		name           string
		format         string
		expectedOutput string
	}{
		{"ShortBash", "short-bash", "TEST_ONE=\"default\"\n"},
		{"LongBash", "long-bash", "\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'.\nTEST_ONE=\"default\"\n"},
		{"ShortDockerfile", "short-dockerfile", "ENV TEST_ONE=\"default\"\n"},
		{"LongDockerfile", "long-dockerfile", "\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'.\nENV TEST_ONE default\n"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			require.NoError(t, env.Print(buffer, tc.format))
			assert.Equal(t, tc.expectedOutput, buffer.String())
		})
	}
}
