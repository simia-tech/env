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

var printFlag *string

// SetUpPrintFlag should be called before a call to `flag.Parse` and sets up an extra flag that is used by
// `EvaluatePrintFlag` to indicate that the environment should be printed.
func SetUpPrintFlag() {
	printFlag = flag.String("print", "", "print the environment in the given format. format can be 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'")
}

// EvaluatePrintFlag tests if the print-flag was given at the program start and prints the registered
// environment fields with thier values to stdout using the specified format. Afterwards, the program exits
// with return code 2. It should be called after the flag has been set up with `SetUpPrintFlag` and the flags
// has been parsed with `flag.Parse`.
func EvaluatePrintFlag() {
	if printFlag == nil {
		return
	}
	if *printFlag != "" {
		if err := Print(os.Stdout, *printFlag); err != nil {
			log.Fatal(err)
		}
		os.Exit(2)
	}
}
