package main

import (
  "fmt"
  "strings"
  "strconv"

  "github.com/aibor/aoc-2021/goutils"
)

func simulate(states [9]int, days int) int {
  sum := 0

  for day := 1; day <= days; day++ {
    wrap := states[0]
    for i := 1; i <= 8; i++ {
      states[i-1] = states[i]
    }

    states[6] += wrap
    states[8] = wrap
  }

  for i := 0; i <= 8; i++ {
    sum += states[i]
  }

  return sum
}

func main() {
  input, _ := goutils.ReadInput()
  states := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

  for _, num := range strings.Split(input[0], ",") {
    numi, _ := strconv.Atoi(num)
    states[numi]++
  }

  fmt.Println("Part 1:", simulate(states, 80))
  fmt.Println("Part 2:", simulate(states, 256))
}
