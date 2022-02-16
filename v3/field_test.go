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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env/v3"
)

func TestField(t *testing.T) {
	t.Run("Bool", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", false)

		t.Run("Value", testSetFn(field, "1", true, nil))
		t.Run("DefaultValue", testUnsetFn(field, false, nil))
		t.Run("ParseError", testSetFn(field, "okaydokay", false, env.ErrInvalidValue))
	})

	t.Run("Int", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", 1)

		t.Run("Value", testSetFn(field, "2", 2, nil))
		t.Run("ParseError", testSetFn(field, "abc", 1, env.ErrInvalidValue))
	})

	t.Run("String", func(t *testing.T) {
		optional := env.Field("OPTIONAL_FIELD", "abc")
		required := env.Field("REQUIRED_FIELD", "abc", env.Required())
		allowed := env.Field("ALLOWED_FIELD", "abc", env.AllowedValues("abc", "def"))

		t.Run("Value", testSetFn(optional, "def", "def", nil))
		t.Run("DefaultValue", testUnsetFn(optional, "abc", nil))
		t.Run("RequiredAndSet", testSetFn(required, "def", "def", nil))
		t.Run("RequiredUnset", testUnsetFn(required, "abc", env.ErrMissingValue))
		t.Run("AllowedValue", testSetFn(allowed, "def", "def", nil))
		t.Run("UnallowedValue", testSetFn(allowed, "ghi", "abc", env.ErrInvalidValue))
	})

	t.Run("StringArray", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", []string{"abc"})

		t.Run("Value", testSetFn(field, "def", []string{"def"}, nil))
		t.Run("DefaultValue", testUnsetFn(field, []string{"abc"}, nil))
	})

	t.Run("StringStringMap", func(t *testing.T) {
		field := env.Field("OPTIONAL_FIELD", map[string]string{"abc": "123"})

		t.Run("Value", testSetFn(field, "def:123", map[string]string{"def": "123"}, nil))
		t.Run("DefaultValue", testUnsetFn(field, map[string]string{"abc": "123"}, nil))
	})
}

func testSetFn[
	T env.FieldType[K, V], K, V env.FieldItemType,
](
	field *env.FieldValue[T, K, V], value string,
	expectValue T, expectErr error,
) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Setenv(field.Name(), value))
		testFn(t, field, expectValue, expectErr)
	}
}

func testUnsetFn[
	T env.FieldType[K, V], K, V env.FieldItemType,
](
	field *env.FieldValue[T, K, V],
	expectValue T, expectErr error,
) func(*testing.T) {
	return func(t *testing.T) {
		require.NoError(t, os.Unsetenv(field.Name()))
		testFn(t, field, expectValue, expectErr)
	}
}

func testFn[
	T env.FieldType[K, V], K, V env.FieldItemType,
](
	tb testing.TB,
	field *env.FieldValue[T, K, V],
	expectValue T, expectErr error,
) {
	value, err := field.Get()
	if expectErr != nil {
		assert.ErrorIs(tb, err, expectErr)
	}
	assert.Equal(tb, value, expectValue)
}
