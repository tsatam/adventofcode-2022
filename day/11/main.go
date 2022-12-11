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
	pt1()
	pt2()
}

func pt1() {
	monkeys := parseInput(input)

	inspected := processRounds(monkeys, 20, false)
	monkeyBusiness := findMonkeyBusiness(inspected)

	fmt.Printf("pt1: [%d]\n", monkeyBusiness)
}
func pt2() {
	monkeys := parseInput(input)

	inspected := processRounds(monkeys, 10000, true)
	monkeyBusiness := findMonkeyBusiness(inspected)

	fmt.Printf("pt2: [%d]\n", monkeyBusiness)
}

func findMonkeyBusiness(inspected []int) int {
	sort.Ints(inspected)

	last := len(inspected) - 1
	return inspected[last] * inspected[last-1]
}
