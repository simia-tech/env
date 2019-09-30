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
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/simia-tech/env"
)

func TestPrint(t *testing.T) {
	env.Clear()
	env.String("TEST_TWO", "default")
	env.String("TEST_ONE", "default", env.Required(), env.AllowedValues("one", "two", "three"))

	testFn := func(format string, expectOutputPattern string) func(*testing.T) {
		return func(t *testing.T) {
			buffer := &bytes.Buffer{}
			require.NoError(t, env.Print(buffer, format))
			assert.Regexp(t, expectOutputPattern, buffer.String())
		}
	}

	t.Run("ShortBash", testFn("short-bash", `^TEST_ONE="default"\nTEST_TWO="default"\n$`))
	t.Run("LongBash", testFn("long-bash", `^\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'. Defined at \S+print_test\.go:\d+\.\nTEST_ONE="default"\n\n# String field. The default value is 'default'. Defined at \S+print_test\.go:\d+\.\nTEST_TWO="default"\n$`))
	t.Run("ShortDockerfile", testFn("short-dockerfile", `^ENV TEST_ONE="default" \\\n    TEST_TWO="default"\n$`))
	t.Run("LongDockerfile", testFn("long-dockerfile", `^\n# String field. Required field. Allowed values are 'one', 'two' and 'three'. The default value is 'default'. Defined at \S+print_test\.go:\d+.\nENV TEST_ONE default\n\n# String field. The default value is 'default'. Defined at \S+print_test\.go:\d+.\nENV TEST_TWO default\n$`))
}
