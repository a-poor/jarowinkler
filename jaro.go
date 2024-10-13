package jarowinkler

import "math"

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func floor(a float64) int {
  return int(math.Floor(a))
}

// Jaro computes the Jaro similarity between two strings.
func Jaro(s1, s2 string) float64 {
  if s1 == s2 {
    return 1.0
  }

  rs1 := []rune(s1)
  rs2 := []rune(s2)

  l1 := len(rs1)
  l2 := len(rs2)

  maxLen := max(l1, l2)
  maxDist := floor(float64(maxLen) / 2.0 - 1.0)

  match := 0

  h1 := make([]int, l1)
  h2 := make([]int, l2)

  for i := 0; i < l1; i++ {
    for j := max(0, i-maxDist); j < min(l2, i+maxDist+1); j++ {
      if rs1[i] == rs2[j] && h2[j] == 0 {
        h1[i] = 1
        h2[j] = 1
        match++
        break
      }
    }
  }
  
  if match == 0 {
    return 0.0
  }

  t := 0.0 // transpositions
  p := 0 // point

  for i := 0; i < l1; i++ {
    if h1[i] == 1 {
      for h2[p] == 0 {
        p++
      }
      
      if rs1[i] != rs2[p] {
        t++
      }
      p++
    }
    t /= 2.0
  }
  
  fmatch := float64(match)
  return (
    fmatch / float64(l1) +
    fmatch / float64(l2) +
    (fmatch - t) /
    float64(match) ) / 3.0
}

// JaroWinkler computes the Jaro-Winkler similarity between two strings.
//
// The Jaro-Winkler similarity is a modification of the Jaro similarity, that
// gives more weight to common prefixes.
//
// The higher the Jaro-Winkler similarity, the more similar the strings are.
// The distance is a number between 0 and 1, where 1 means exact match and
// 0 means no similarity.
func JaroWinkler(s1, s2 string) float64 {
  j := Jaro(s1, s2)
  if j < 0.7 {
    return j
  }

  rs1 := []rune(s1)
  rs2 := []rune(s2)

  l1 := len(rs1)
  l2 := len(rs2)
  prefix := 0

  for i := 0; i < min(l1, l2); i++ {
    if rs1[i] == rs2[i] {
      prefix++
    } else {
      break
    }
  }
  prefix = min(prefix, 4)

  return j + float64(prefix) * 0.1 * (1.0 - j)
}

