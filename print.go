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
	"sort"
)

// Print prints the environment in the provided format.
func Print(w io.Writer, format string) error {
	eh := ErrorHandler
	ErrorHandler = NullErrorHandler
	defer func() {
		ErrorHandler = eh
	}()

	switch format {
	case "short-bash":
		printShortBash(w)
		return nil
	case "long-bash":
		printLongBash(w)
		return nil
	case "short-dockerfile":
		printShortDockerfile(w)
		return nil
	case "long-dockerfile":
		printLongDockerfile(w)
		return nil
	default:
		return fmt.Errorf("unknown format '%s'. known values are 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'", format)
	}
}

func printShortBash(w io.Writer) {
	forEachField(func(field Field) {
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	})
}

func printLongBash(w io.Writer) {
	forEachField(func(field Field) {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	})
}

func printShortDockerfile(w io.Writer) {
	index := 0
	forEachField(func(field Field) {
		if index == 0 {
			fmt.Fprintf(w, "ENV ")
		} else {
			fmt.Fprintf(w, " \\\n    ")
		}
		fmt.Fprintf(w, "%s=\"%s\"", field.Name(), field.Value())
		index++
	})
	fmt.Fprintln(w)
}

func printLongDockerfile(w io.Writer) {
	forEachField(func(field Field) {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "ENV %s %s\n", field.Name(), field.Value())
	})
}

func forEachField(fn func(Field)) {
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fn(fields[key])
	}
}
