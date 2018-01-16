package env_test

import (
	"bytes"
	"testing"

	"github.com/simia-tech/env"
	"github.com/stretchr/testify/assert"
)

func TestPrintBashWithGeneratedDescription(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Required(), env.AllowedValues("one", "two", "three"))

	buffer := &bytes.Buffer{}
	env.PrintBash(buffer)

	assert.Equal(t, "\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'.\nTEST_ONE=\"default\"\n", buffer.String())
}

func TestPrintBashWithCustomDescription(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Description("Test field."))

	buffer := &bytes.Buffer{}
	env.PrintBash(buffer)

	assert.Equal(t, "\n# Test field.\nTEST_ONE=\"default\"\n", buffer.String())
}

func TestPrintDockerfileWithGeneratedDescription(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Required(), env.AllowedValues("one", "two", "three"))

	buffer := &bytes.Buffer{}
	env.PrintDockerfile(buffer)

	assert.Equal(t, "\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'.\nENV TEST_ONE default\n", buffer.String())
}
