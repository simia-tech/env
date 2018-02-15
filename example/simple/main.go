package main

import (
	"flag"
	"fmt"

	"github.com/simia-tech/env"
)

var (
	name = env.String("NAME", "simia")
	age  = env.Int("AGE", 18)
)

func main() {
	env.SetUpPrintFlag()
	flag.Parse()
	env.EvaluatePrintFlag()

	fmt.Printf("%s is %d years old\n", name.Get(), age.Get())
}
