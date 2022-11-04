package main

import (
  "bufio"
  "fmt"
  "os"
  "errors"
  "sort"
)

var charmap = map[string]string{
  "(": ")",
  "[": "]",
  "{": "}",
  "<": ">",
}

var p1_scoremap = map[string]int{
  ")": 3,
  "]": 57,
  "}": 1197,
  ">": 25137,
}
var p2_scoremap = map[string]int{
  ")": 1,
  "]": 2,
  "}": 3,
  ">": 4,
}


func Parse(line string) (string, int, error) {
  stack := [128]string{}
  stackidx := -1

  for strpos, char := range line {
    str := string(char)
    _, ok := charmap[str]

    if ok == true {
      stackidx++
      stack[stackidx] = str
      continue
    }

    lastopen := stack[stackidx]
    closechar := charmap[lastopen]

    if str == closechar {
      stackidx--
      continue
    }
    return str, strpos, errors.New("unexpected char")
  }

  retstr := ""
  for i := stackidx; i >= 0; i-- {
    retstr += charmap[stack[i]]
  }

  return retstr, len(line), nil
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  p1_score := 0
  p2_scores := []int{}

  for scanner.Scan() {
    str, _, err := Parse(scanner.Text())
    if err != nil {
      p1_score += p1_scoremap[str]
    } else {
      line_score := 0
      for _, char := range str {
        line_score = line_score * 5 + p2_scoremap[string(char)]
      }
      p2_scores = append(p2_scores, line_score)
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  p2_scoresSlice := p2_scores[:]
  sort.Ints(p2_scoresSlice)

  fmt.Println("Part 1:", p1_score)
  fmt.Println("Part 2:", p2_scoresSlice[int(len(p2_scoresSlice)/2)])
}

