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

func NewStringFieldsIterator(s string) StringFieldIterator {
	return StringFieldIterator{NewIterator(strings.Fields(s))}
}

type StringFieldIterator struct {
	Iterator[string]
}

func (i *StringFieldIterator) Int() (int, error) {
	return strconv.Atoi(i.Value())
}

func (i *StringFieldIterator) MustBeInt() int {
	return MustBeInt(i.Value())
}

func MustBeInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("must parse Atoi")
	}
	return n
}
