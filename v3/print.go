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
	"fmt"
	"io"
)

// Print prints the environment in the provided format.
func Print(w io.Writer, format string) error {
	p, ok := printer[format]
	if !ok {
		return fmt.Errorf("unknown format '%s'. known values are 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'", format)
	}
	p(w)
	return nil
}

var printer = map[string]func(io.Writer){
	"short-bash":       printShortBash,
	"long-bash":        printLongBash,
	"short-dockerfile": printShortDockerfile,
	"long-dockerfile":  printLongDockerfile,
}

func printShortBash(w io.Writer) {
	for _, field := range fields {
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.GetRawOrDefault())
	}
}

func printLongBash(w io.Writer) {
	for _, field := range fields {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.GetRawOrDefault())
	}
}

func printShortDockerfile(w io.Writer) {
	index := 0
	for _, field := range fields {
		if index == 0 {
			fmt.Fprintf(w, "ENV ")
		} else {
			fmt.Fprintf(w, " \\\n    ")
		}
		fmt.Fprintf(w, "%s=\"%s\"", field.Name(), field.GetRawOrDefault())
		index++
	}
	fmt.Fprintln(w)
}

func printLongDockerfile(w io.Writer) {
	for _, field := range fields {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "ENV %s %s\n", field.Name(), field.GetRawOrDefault())
	}
}
