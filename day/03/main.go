package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("%d\n", processInput(input))
}

func processInput(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		sum += processLine(line)
	}

	return sum
}

func processLine(line string) int {
	size := len(line) / 2

	first, second := line[:size], line[size:]

	for idx, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		if strings.ContainsRune(first, char) && strings.ContainsRune(second, char) {
			return idx + 1
		}
	}

	return 0
}
