package env

import (
	"fmt"
	"io"
)

// Print writes an example bash script with all registered environment field and it's default values
// to the provided io.Writer.
func Print(w io.Writer) {
	eh := ErrorHandler
	ErrorHandler = NullErrorHandler
	for _, field := range fields {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	}
	ErrorHandler = eh
}
