package goutils

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput() (input []string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	err = scanner.Err()
	return
}

func SplitInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
