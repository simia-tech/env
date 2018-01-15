package env_test

import (
	"bytes"
	"testing"

	"github.com/simia-tech/env"
	"github.com/stretchr/testify/assert"
)

func TestPrintGeneratedDescription(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Required(), env.AllowedValues("one", "two", "three"))

	buffer := &bytes.Buffer{}
	env.Print(buffer)

	assert.Equal(t, "\n# Required field. Allowed values are 'one', 'two' and 'three'.\nTEST_ONE=\"default\"\n", buffer.String())
}

func TestPrintCustomDescription(t *testing.T) {
	env.Clear()
	env.String("TEST_ONE", "default", env.Description("Test field."))

	buffer := &bytes.Buffer{}
	env.Print(buffer)

	assert.Equal(t, "\n# Test field.\nTEST_ONE=\"default\"\n", buffer.String())
}
