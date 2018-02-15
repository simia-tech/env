
# env

[![GoDoc](https://godoc.org/github.com/simia-tech/env?status.svg)](https://godoc.org/github.com/simia-tech/env) [![Build Status](https://travis-ci.org/simia-tech/env.svg?branch=master)](https://travis-ci.org/simia-tech/env)

Golang handling of environment values

## Example

```go
var (
	name = env.String("NAME", "joe")
	age  = env.Int("AGE", 24)
)

func main() {
	env.SetUpPrintFlag()
	flag.Parse()
	env.EvaluatePrintFlag()

	fmt.Printf("%s is %d years old\n", name.Get(), age.Get())
}
```

If the program is called with `-print short-bash`, all registered environment fields would be printed...

```bash
NAME="joe"
AGE="24"
```

By using `-print long-bash`, a description for each field is generated.

```bash
# String field. The default value is 'joe'.
NAME="joe"

# Int field. The default value is '24'.
AGE="24"
```

## License

The project is licensed under [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0).
