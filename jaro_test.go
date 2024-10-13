package jarowinkler_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/a-poor/jarowinkler"
)

func roundTo(n float64, places int) float64 {
  multiplier := math.Pow(10, float64(places))
  return math.Round(n*multiplier) / multiplier
}

func TestJaro(t *testing.T) {
  round := 6
  cases := []struct {
    s1, s2 string
    want float64
  }{
    {"", "", 1.0},
    {"CRATE", "TRACE", 0.733333},
  }

  for i, c := range cases {
    got := jarowinkler.Jaro(c.s1, c.s2)
    ground := roundTo(got, round)
    if ground != c.want {
      t.Errorf("[%02d] Jaro(%q, %q) == %f, want %f", i, c.s1, c.s2, ground, c.want)
    }
  }
}


func TestJaroWinkler(t *testing.T) {
  round := 4
  cases := []struct {
    s1, s2 string
    want float64
  }{
    {"", "", 1.0},
    {"aaa", "aaa", 1.0},
    {"aaa", "bbb", 0.0},
    {"TRATE", "TRACE", 0.9067},
    {"apple", "applet", 0.9667},
    {
      "10000000000000000000000000000000000000000000000000000000000000020",
      "00000000000000000000000000000000000000000000000000000000000000000",
      0.9795,
    },
    {
      "00000000000000100000000000000000000000010000000000000000000000000",
      "0000000000000000000000000000000000000000000000000000000000000000000000000000001",
      0.9533,
    },
    {
      "00000000000000000000000000000000000000000000000000000000000000000",
      "01000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      0.8523,
    },
  }

  for i, c := range cases {
    // Try forward
    got := jarowinkler.JaroWinkler(c.s1, c.s2)
    ground := roundTo(got, round)
    if ground != c.want {
      t.Errorf("[%02d] JaroWinkler(%q, %q) == %f, want %f", i, c.s1, c.s2, ground, c.want)
    }
   
    // Try reverse
    got = jarowinkler.JaroWinkler(c.s2, c.s1)
    ground = roundTo(got, round)
    if ground != c.want {
      t.Errorf("[%02d] JaroWinkler(%q, %q) == %f, want %f", i, c.s2, c.s1, ground, c.want)
    }
  }
}

func BenchmarkJaro(b *testing.B) {
  s1 := "CRATE"
  s2 := "CRACE"
  for i := 0; i < b.N; i++ {
    jarowinkler.Jaro(s1, s2)
  }
}

func BenchmarkJaroWinkler(b *testing.B) {
  s1 := "CRATE"
  s2 := "CRACE"
  for i := 0; i < b.N; i++ {
    jarowinkler.JaroWinkler(s1, s2)
  }
}

func BenchmarkJaroWinklerLong(b *testing.B) {
  s1 := "00000000000000000000000000000000000000000000000000000000000000000"
  s2 := "01000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
  for i := 0; i < b.N; i++ {
    jarowinkler.JaroWinkler(s1, s2)
  }
}

func ExampleJaro() {
  s1 := "CRATE"
  s2 := "CRACE"
  fmt.Printf("%.4f\n", jarowinkler.Jaro(s1, s2))
  // Output: 0.8667
}

func ExampleJaroWinkler() {
  s1 := "CRATE"
  s2 := "CRACE"
  fmt.Printf("%.4f\n", jarowinkler.JaroWinkler(s1, s2))
  // Output: 0.9067
}




