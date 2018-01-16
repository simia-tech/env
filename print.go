package env

import (
	"fmt"
	"io"
)

// PrintBash writes an example bash script with all registered environment field and it's values
// to the provided io.Writer.
func PrintBash(w io.Writer) {
	eh := ErrorHandler
	ErrorHandler = NullErrorHandler
	for _, field := range fields {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	}
	ErrorHandler = eh
}

// PrintDockerfile writes an example dockerfile with all registered environment field and it's values
// to the provided io.Writer.
func PrintDockerfile(w io.Writer) {
	eh := ErrorHandler
	ErrorHandler = NullErrorHandler
	for _, field := range fields {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "ENV %s %s\n", field.Name(), field.Value())
	}
	ErrorHandler = eh
}
