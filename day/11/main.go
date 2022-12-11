package main

import (
	_ "embed"
	"fmt"
	"sort"
)

var (
	//go:embed input
	input string
)

func main() {
	monkeys := parseInput(input)
	monkeys = processRounds(monkeys, 20)
	monkeyBusiness := findMonkeyBusiness(monkeys)

	fmt.Printf("pt1: [%d]\n", monkeyBusiness)
}

func findMonkeyBusiness(monkeys []Monkey) int {
	inspected := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspected[i] = monkey.inspected
	}

	sort.Ints(inspected)

	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}
