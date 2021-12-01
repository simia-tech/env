
# env

[![GoDoc](https://godoc.org/github.com/simia-tech/env?status.svg)](https://godoc.org/github.com/simia-tech/env) [![Build Status](https://travis-ci.org/simia-tech/env.svg?branch=master)](https://travis-ci.org/simia-tech/env)

Implements a simple way of handling environment values. Each environment field is simply reflected by a
variable inside the Go program. Out of the box handlers for the types `bool`, `[]byte`, `time.Duration`,
`int`, `[]int`, `string` and `[]string` are provided. Other types can be added by using
the `RegisterField` function.

## Example

```go
var (
    name = env.String("NAME", "joe")
    age  = env.Int("AGE", 24)
)

func main() {
    env.ParseFlags()

    fmt.Printf("%s is %d years old\n", name.Get(), age.Get())
}
```

If the program is called with `-print-env`, all registered environment fields would be printed...

```bash
NAME="joe"
AGE="24"
```

By using `-print-env -print-env-format long-bash`, a description for each field is generated.

```bash
# String field. The default value is 'joe'. Defined at .../main.go:11.
NAME="joe"

# Int field. The default value is '24'. Defined at .../main.go:10.
AGE="24"
```

## License

The project is licensed under [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0).
