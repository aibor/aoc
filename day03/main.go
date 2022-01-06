package main

import (
  "fmt"
  "math"
  "strings"
  "strconv"

  "github.com/aibor/aoc-2021/goutils"
)

func findnums(data []string, pos int, maj bool) string {
  new_arr := [2][]string{[]string{},[]string{}}
  major := 1
  if maj == false {
    major = 0
  }

  if len(data) == 1 {
    return data[0]
  }

  for _, num := range data {
    idx := 0
    if string(num[pos]) == "1" {
      idx = 1
    }
    new_arr[idx] = append(new_arr[idx], num)
  }

  if len(new_arr[0]) > len(new_arr[1]) {
    if maj {
      major = 0
    } else {
      major = 1
    }
  }

  return findnums(new_arr[major], pos + 1, maj)
}

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

  oxygen, _ := strconv.ParseInt(findnums(input, 0, true), 2, 32)
  carbon, _ := strconv.ParseInt(findnums(input, 0, false), 2, 32)

  fmt.Println("Part 2:", oxygen * carbon)
}
