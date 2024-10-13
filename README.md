# JaroWinkler

[![Go Reference](https://pkg.go.dev/badge/github.com/a-poor/jarowinkler.svg)](https://pkg.go.dev/github.com/a-poor/jarowinkler)
[![Go Test](https://github.com/a-poor/jarowinkler/actions/workflows/test.yaml/badge.svg)](https://github.com/a-poor/jarowinkler/actions/workflows/test.yaml)


An implementation of the Jaro-Winkler string similarity algorithm in Go.

## Installation

```sh
go get github.com/a-poor/jarowinkler
```

## Usage

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

## Benchmarks

There are (very basic) benchmarks for the `Jaro` and `JaroWinkler` functions,
though they don't compare the performance of this implementation to other text
similarity algorithms.

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/a-poor/jarowinkler
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
BenchmarkJaro-8              	 2647617	       406.6 ns/op
BenchmarkJaroWinkler-8       	 2698708	       434.8 ns/op
BenchmarkJaroWinklerLong-8   	  111476	     10478 ns/op
PASS
ok  	github.com/a-poor/jarowinkler	4.454s
```


