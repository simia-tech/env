package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/v3/internal/parser"
)

func TestParseInts(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		testFn := func(raw string, expected []int) (string, func(*testing.T)) {
			return raw, func(t *testing.T) {
				values, err := parser.ParseInts(raw)
				require.NoError(t, err)
				assert.Equal(t, expected, values)
			}
		}

		t.Run(testFn(``, []int{}))
		t.Run(testFn(`1`, []int{1}))
		t.Run(testFn(`1,2`, []int{1, 2}))
		t.Run(testFn(`1, 2`, []int{1, 2}))
		t.Run(testFn(`"1", 2`, []int{1, 2}))
	})

	t.Run("Format", func(t *testing.T) {
		testFn := func(raw []int, expected string) (string, func(*testing.T)) {
			return expected, func(t *testing.T) {
				assert.Equal(t, expected, parser.FormatInts(raw))
			}
		}

		t.Run(testFn([]int{}, ""))
		t.Run(testFn([]int{1}, "1"))
		t.Run(testFn([]int{1, 2}, "1,2"))
	})
}
