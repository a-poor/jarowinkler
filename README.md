# JaroWinkler

[![Go Reference](https://pkg.go.dev/badge/github.com/a-poor/jarowinkler.svg)](https://pkg.go.dev/github.com/a-poor/jarowinkler)
[![Go Test](https://github.com/a-poor/jarowinkler/actions/workflows/test.yaml/badge.svg)](https://github.com/a-poor/jarowinkler/actions/workflows/test.yaml)


An implementation of the Jaro-Winkler string similarity algorithm in Go.

## Usage

```sh
go get github.com/a-poor/jarowinkler
```

```go
package main

import (
    "fmt"

    "github.com/a-poor/jarowinkler"
)

func main() {
  s1 := "CRATE"
  s2 := "CRACE"
  fmt.Printf("%.4f\n", jarowinkler.JaroWinkler(s1, s2))
  // Output: 0.9067
}

```

## References

Based on the following implementations:

- https://geeksforgeeks.org/jaro-and-jaro-winkler-similarity/
- https://en.wikipedia.org/wiki/Jaro-Winkler_distance
- https://safjan.com/jaro-winkler-similarity
- https://tech.popdata.org/speeding-up-Jaro-Winkler-with-rust-and-bitwise-operations

