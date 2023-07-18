// Copyright 2018 Philipp Brüll <pb@simia.tech>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/v3"
)

func TestField(t *testing.T) {
	t.Run("Bool", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", false)

		t.Run("Value", testSetFn(field, "1", true, nil))
		t.Run("DefaultValue", testUnsetFn(field, false, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, "false", nil))
		t.Run("ParseError", testSetFn(field, "okaydokay", false, env.ErrInvalidValue))
	})

	t.Run("Bytes", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", []byte{0, 1, 2, 3})

		t.Run("Value", testSetFn(field, "ffeeddcc", []byte{0xff, 0xee, 0xdd, 0xcc}, nil))
		t.Run("DefaultValue", testUnsetFn(field, []byte{0, 1, 2, 3}, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, "00010203", nil))
		t.Run("ParseError", testSetFn(field, "okaydokay", []byte{0, 1, 2, 3}, env.ErrInvalidValue))
	})

	t.Run("Duration", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", 5*time.Second)

		t.Run("Value", testSetFn(field, "10s", 10*time.Second, nil))
		t.Run("DefaultValue", testUnsetFn(field, 5*time.Second, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, "5s", nil))
		t.Run("ParseError", testSetFn(field, "okaydokay", 5*time.Second, env.ErrInvalidValue))
	})

	t.Run("Int", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", 1)

		t.Run("Value", testSetFn(field, "2", 2, nil))
		t.Run("DefaultValue", testUnsetFn(field, 1, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, "1", nil))
		t.Run("ParseError", testSetFn(field, "abc", 1, env.ErrInvalidValue))
	})

	t.Run("IntArray", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", []int{1})

		t.Run("Value", testSetFn(field, "2", []int{2}, nil))
		t.Run("DefaultValue", testUnsetFn(field, []int{1}, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, "1", nil))
	})

	t.Run("String", func(t *testing.T) {
		optional := env.Field("OPTIONAL_FIELD", "abc")
		required := env.Field("REQUIRED_FIELD", "abc", env.Required())
		allowed := env.Field("ALLOWED_FIELD", "abc", env.AllowedValues("abc", "def"))

		t.Run("Value", testSetFn(optional, "def", "def", nil))
		t.Run("DefaultValue", testUnsetFn(optional, "abc", nil))
		t.Run("RawDefaultValue", testRawUnsetFn(optional, "abc", nil))
		t.Run("RequiredAndSet", testSetFn(required, "def", "def", nil))
		t.Run("RequiredUnset", testUnsetFn(required, "abc", env.ErrMissingValue))
		t.Run("AllowedValue", testSetFn(allowed, "def", "def", nil))
		t.Run("UnallowedValue", testSetFn(allowed, "ghi", "abc", env.ErrInvalidValue))
	})

	t.Run("StringArray", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", []string{"abc"})

		t.Run("Value", testSetFn(field, "def", []string{"def"}, nil))
		t.Run("DefaultValue", testUnsetFn(field, []string{"abc"}, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, `"abc"`, nil))
	})

	t.Run("StringStringMap", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", map[string]string{"abc": "123"})

		t.Run("Value", testSetFn(field, "def:123", map[string]string{"def": "123"}, nil))
		t.Run("DefaultValue", testUnsetFn(field, map[string]string{"abc": "123"}, nil))
		t.Run("RawDefaultValue", testRawUnsetFn(field, `abc:"123"`, nil))
	})
}

func testSetFn[T env.FieldType](field *env.F[T], value string, expectValue T, expectErr error) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Setenv(field.Name(), value))
		testFn(t, field, expectValue, expectErr)
	}
}

func testUnsetFn[T env.FieldType](field *env.F[T], expectValue T, expectErr error) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Unsetenv(field.Name()))
		testFn(t, field, expectValue, expectErr)
	}
}

func testFn[T env.FieldType](tb testing.TB, field *env.F[T], expectValue T, expectErr error) {
	value, err := field.Get()
	if expectErr != nil {
		assert.ErrorIs(tb, err, expectErr)
	}
	assert.Equal(tb, expectValue, value)
}

func testRawSetFn[T env.FieldType](field *env.F[T], value string, expectValue string, expectErr error) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Setenv(field.Name(), value))
		testRawFn(t, field, expectValue, expectErr)
	}
}

func testRawUnsetFn[T env.FieldType](field *env.F[T], expectValue string, expectErr error) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Unsetenv(field.Name()))
		testRawFn(t, field, expectValue, expectErr)
	}
}

func testRawFn[T env.FieldType](tb testing.TB, field *env.F[T], expectValue string, expectErr error) {
	value, err := field.GetRaw()
	if expectErr != nil {
		assert.ErrorIs(tb, err, expectErr)
	}
	assert.Equal(tb, expectValue, value)
}
