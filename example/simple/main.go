package main

import (
	"fmt"

	"github.com/simia-tech/env"
)

var (
	name   = env.String("NAME", "joe")
	age    = env.Int("AGE", 24)
	shifts = env.StringMap("SHIFTS", map[string]string{"monday": "9am - 5pm"})
)

func main() {
	env.ParseFlags()

	fmt.Printf("%s is %d years old\nshifts are %v\n",
		name.GetOrDefault(), age.GetOrDefault(), shifts.GetOrDefault())
}
