
# env

[![GoDoc](https://godoc.org/github.com/simia-tech/env?status.svg)](https://godoc.org/github.com/simia-tech/env) [![Build Status](https://travis-ci.org/simia-tech/env.svg?branch=master)](https://travis-ci.org/simia-tech/env)

Golang handling of environment values

## Example

```go
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
```

## License

The project is licensed under [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0).
