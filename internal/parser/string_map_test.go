package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/internal/parser"
)

func TestParseStringMap(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		testFn := func(raw string, expected map[string]string) (string, func(*testing.T)) {
			return raw, func(t *testing.T) {
				m, err := parser.ParseStringMap(raw)
				require.NoError(t, err)
				assert.Equal(t, expected, m)
			}
		}

		t.Run(testFn(``, map[string]string{}))
		t.Run(testFn(`one`, map[string]string{"one": ""}))
		t.Run(testFn(`one:value`, map[string]string{"one": "value"}))
		t.Run(testFn(`one:value 123`, map[string]string{"one": "value 123"}))
		t.Run(testFn(`one:value,two`, map[string]string{"one": "value", "two": ""}))
		t.Run(testFn(`one:"value"`, map[string]string{"one": "value"}))
		t.Run(testFn(`one:'value'`, map[string]string{"one": "value"}))
		t.Run(testFn(`one:"value \"123\""`, map[string]string{"one": `value "123"`}))
		t.Run(testFn(`one:'value \'123\''`, map[string]string{"one": `value '123'`}))
		t.Run(testFn(`one:'value "123"'`, map[string]string{"one": `value "123"`}))
		t.Run(testFn(`one:"value '123'"`, map[string]string{"one": `value '123'`}))
	})

	t.Run("Format", func(t *testing.T) {
		testFn := func(raw map[string]string, expected string) (string, func(*testing.T)) {
			return expected, func(t *testing.T) {
				assert.Equal(t, expected, parser.FormatStringMap(raw))
			}
		}

		t.Run(testFn(map[string]string{}, ""))
		t.Run(testFn(map[string]string{"one": ""}, "one"))
		t.Run(testFn(map[string]string{"one": "", "two": ""}, "one,two"))
		t.Run(testFn(map[string]string{"one": "value"}, `one:"value"`))
		t.Run(testFn(map[string]string{"one": "value", "two": "value"}, `one:"value",two:"value"`))
		t.Run(testFn(map[string]string{"one": `value "123"`}, `one:"value \"123\""`))
	})
}
