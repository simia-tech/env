package main

import (
	"flag"
	"fmt"

	"github.com/simia-tech/env"
)

var (
	name = env.String("NAME", "joe")
	age  = env.Int("AGE", 24)
)

func main() {
	env.SetUpFlags()
	flag.Parse()
	env.EvaluateFlags()

	fmt.Printf("%s is %d years old\n", name.Get(), age.Get())
}
