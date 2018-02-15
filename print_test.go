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
