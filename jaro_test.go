package jarowinkler_test

import (
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
  round := 6
  cases := []struct {
    s1, s2 string
    want float64
  }{
    {"", "", 1.0},
    {"TRATE", "TRACE", 0.906667},
    {"apple", "applet", 0.966667},
  }

  for i, c := range cases {
    got := jarowinkler.JaroWinkler(c.s1, c.s2)
    ground := roundTo(got, round)
    if ground != c.want {
      t.Errorf("[%02d] JaroWinkler(%q, %q) == %f, want %f", i, c.s1, c.s2, ground, c.want)
    }
  }
}

