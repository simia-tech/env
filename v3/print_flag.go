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

package env

import (
	"flag"
	"log"
	"os"
)

// ParseFlags tests if the print-flag was given at the program start and prints the registered
// environment fields with thier values to stdout using the specified format. Afterwards, the program exits
// with return code 2.
func ParseFlags() {
	WithFlags(func() {
		flag.Parse()
	})
}

func WithFlags(fn func()) {
	printEnvFlag := flag.Bool("print-env", false, "print the environment with the current values")
	printEnvFormatFlag := flag.String("print-env-format", "short-bash", "print the environment in the given format. format can be 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'")

	fn()

	if *printEnvFlag {
		if err := Print(os.Stdout, *printEnvFormatFlag); err != nil {
			log.Fatal(err)
		}
		os.Exit(2)
	}
}
