CAST
====
![Travis CI Build Status](https://api.travis-ci.org/tidwall/cast.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/tidwall/cast?status.svg)](https://godoc.org/github.com/tidwall/cast)

Quickly convert string <-> []byte without memory reallocations and create mutable string or immutable []byte.

This package is a **danger zone** and should not be entered without understanding the ground rules.

1. Converting a string -> []byte will result in an immutable byte slice. Editing will cause a panic.
2. Converting a []byte -> string will result in a mutable string. Editing the originial bytes will change the string too.


Create immutable []byte:

```go
var s = "Hello Planet"
b := cast.ToBytes(s)
fmt.Printf("%s\n", "J"+string(b[1:]))

// Output:
// Jello Planet
```

Create mutable string:

```go
var b = []byte("Hello Planet")
s := cast.ToString(b)
b[0] = 'J'
fmt.Printf("%s\n", s)

// Output:
// Jello Planet
```

## Contact
Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

CAST source code is available under the MIT [License](/LICENSE).
