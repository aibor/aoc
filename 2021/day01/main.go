package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var slidwin [4]int

func slidwincomp(dec int, i int) bool {
	return i > dec && slidwin[i%4] > slidwin[(i-dec)%4]
}

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var part_one_inc, part_two_inc, cur, i int = 0, 0, 0, 0

	for scanner.Scan() {
		cur, _ = strconv.Atoi(scanner.Text())
		i++

		slidwin[i%4] = cur

		if slidwincomp(1, i) {
			part_one_inc++
		}

		if slidwincomp(3, i) {
			part_two_inc++
		}
	}

	fmt.Println("Part 1:", part_one_inc)
	fmt.Println("Part 2:", part_two_inc)
}
