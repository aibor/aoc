package main

import (
  "fmt"
  "math"
  "strings"
  "strconv"

  "github.com/aibor/aoc-2021/goutils"
)

func main() {
  input, _ := goutils.ReadInput()
  var gamma, epsilon int

  bits := make(map[int]int)
  lines := len(input)

  for _, line := range input {
    nums := strings.Split(line, "")
    c := len(nums)
    for n, num := range nums {
      numi, _ := strconv.Atoi(num)
      bits[c - n - 1] += numi
    }
  }

  for key, num := range bits {
    if num > lines/2 {
      gamma += int(math.Pow(float64(2), float64(key)))
    }
    if num < lines/2 {
      epsilon += int(math.Pow(float64(2), float64(key)))
    }
  }

  fmt.Println("Part 1:", gamma * epsilon)
  //fmt.Println("Part 2:", )
}
