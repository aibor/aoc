package goutils

import (
	"bufio"
	"os"
	"strconv"
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

func NewStringFieldsIterator(s string) Iterator[string] {
	return NewIterator(strings.Fields(s))
}

func MustBeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("must parse Atoi")
	}
	return n
}
