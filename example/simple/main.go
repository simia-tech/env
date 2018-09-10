package main

import (
	"fmt"

	"github.com/simia-tech/env"
)

var (
	name = env.String("NAME", "joe")
	age  = env.Int("AGE", 24)
)

func main() {
	env.ParseFlags()

	fmt.Printf("%s is %d years old\n", name.Get(), age.Get())
}
