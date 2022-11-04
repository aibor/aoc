package main

import (
  "fmt"
  "strings"
  "strconv"

  "github.com/aibor/aoc/goutils"
)

func main() {
  input, _ := goutils.ReadInput()
  var fwd, aim, depth int
  mv := make(map[string]int)

  for _, line := range input {
    inst := strings.Split(line, " ")
    dir := inst[0]
    dist, _ := strconv.Atoi(inst[1])

    switch dir {
      case "down":
        aim += dist
      case "up":
        aim -= dist
      case "forward":
        fwd += dist
        depth += dist * aim
    }

    mv[dir] += dist
  }

  fmt.Println("Part 1:", mv["forward"] * (mv["down"] - mv["up"]))
  fmt.Println("Part 2:", fwd * depth)
}
