package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func main() {
  file, err := os.Open("input")

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var prev, cur, inc int = 0, 0, 0

  for scanner.Scan() {
    cur, _ = strconv.Atoi(scanner.Text())

    if prev != 0 && cur > prev {
      inc++
    }

    prev = cur
  }

  fmt.Println(inc)
}
