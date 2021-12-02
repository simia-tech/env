package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/internal/parser"
)

func TestParseStrings(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		testFn := func(raw string, expected []string) (string, func(*testing.T)) {
			return raw, func(t *testing.T) {
				values, err := parser.ParseStrings(raw)
				require.NoError(t, err)
				assert.Equal(t, expected, values)
			}
		}

		t.Run(testFn(``, []string{}))
		t.Run(testFn(`one`, []string{"one"}))
		t.Run(testFn(`one,two`, []string{"one", "two"}))
		t.Run(testFn(`one, two`, []string{"one", "two"}))
		t.Run(testFn(`one,"two"`, []string{"one", "two"}))
		t.Run(testFn(`one,"two \"123\""`, []string{"one", `two "123"`}))
	})

	t.Run("Format", func(t *testing.T) {
		testFn := func(raw []string, expected string) (string, func(*testing.T)) {
			return expected, func(t *testing.T) {
				assert.Equal(t, expected, parser.FormatStrings(raw))
			}
		}

		t.Run(testFn([]string{}, ""))
		t.Run(testFn([]string{"one"}, `"one"`))
		t.Run(testFn([]string{"one", "two"}, `"one","two"`))
		t.Run(testFn([]string{"one", `two "123"`}, `"one","two \"123\""`))
	})
}
