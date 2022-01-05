package goutils

import (
  "bufio"
  "os"
)

func ReadInput() (input []string, err error) {
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    input = append(input, scanner.Text())
  }

  err = scanner.Err()
  return
}
