package env

import (
	"flag"
	"fmt"
	"io"
)

// PrintFlag returns a flag for the environment print.
func PrintFlag() *string {
	return flag.String("print", "", "print the environment in the given format. format can be 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'")
}

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
	for _, field := range fields {
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	}
}

func printLongBash(w io.Writer) {
	for _, field := range fields {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "%s=\"%s\"\n", field.Name(), field.Value())
	}
}

func printShortDockerfile(w io.Writer) {
	for index, field := range fields {
		if index == 0 {
			fmt.Fprintf(w, "ENV ")
		} else {
			fmt.Fprintf(w, " \\\n    ")
		}
		fmt.Fprintf(w, "%s=\"%s\"", field.Name(), field.Value())
	}
	fmt.Fprintln(w)
}

func printLongDockerfile(w io.Writer) {
	for _, field := range fields {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "# %s\n", field.Description())
		fmt.Fprintf(w, "ENV %s %s\n", field.Name(), field.Value())
	}
}
