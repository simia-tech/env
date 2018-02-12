package env

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var printFlag = flag.String("print", "", "print the environment in the given format. format can be 'short-bash', 'long-bash', 'short-dockerfile' and 'long-dockerfile'")

// EvaluatePrintFlag tests if the print-flag was given at the program start and prints the registered
// environment fields with thier values to stdout using the specified format. Afterwards, the program exits
// with return code 2.
func EvaluatePrintFlag() {
	if *printFlag != "" {
		if err := Print(os.Stdout, *printFlag); err != nil {
			log.Fatal(err)
		}
		os.Exit(2)
	}
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
