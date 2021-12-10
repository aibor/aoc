package main

import (
  "bufio"
  "fmt"
  "os"
  "errors"
)

var charmap = map[string]string{
  "(": ")",
  "[": "]",
  "{": "}",
  "<": ">",
}

var scoremap = map[string]int{
  ")": 3,
  "]": 57,
  "}": 1197,
  ">": 25137,
}

func Parse(line string) (string, int, error) {
  var stack[128] string
  var stackidx int = -1
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
  return "", len(line), nil
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  score := 0

  for scanner.Scan() {
    char, _, err := Parse(scanner.Text())
    if err != nil {
      score += scoremap[char]
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  fmt.Println("Part 1:", score)
}

