package main

import (
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) []Monkey {
	rawMonkeys := strings.Split(strings.TrimSpace(input), "\n\n")

	monkeys := make([]Monkey, len(rawMonkeys))

	for i, rawMonkey := range rawMonkeys {
		monkeys[i] = parseMonkey(rawMonkey)
	}

	return monkeys
}

func parseMonkey(rawMonkey string) Monkey {
	lines := strings.Split(rawMonkey, "\n")

	return Monkey{
		items:      parseStartingItems(lines[1]),
		operation:  parseOperation(lines[2]),
		testDiv:    parseTestDiv(lines[3]),
		throwTrue:  parseThrowTrue(lines[4]),
		throwFalse: parseThrowFalse(lines[5]),

		inspected: 0,
	}
}

func parseStartingItems(line string) []int {
	rawItems := strings.Split(line[18:], ", ")
	items := make([]int, len(rawItems))

	for i, rawItem := range rawItems {
		item, _ := strconv.Atoi(rawItem)
		items[i] = int(item)
	}

	return items
}

func parseOperation(line string) Operation {
	rawOperation := line[19:]

	if rawOperation == "old * old" {
		return PowOperation{}
	}

	value := parseSingleIntAtPosition(rawOperation, 6)
	if rawOperation[4] == '+' {
		return AddOperation{int(value)}
	}
	if rawOperation[4] == '*' {
		return MulOperation{int(value)}
	}

	log.Fatalf("Could not parse operation: [%s]", line)
	return nil
}

func parseTestDiv(line string) int {
	return parseSingleIntAtPosition(line, 21)
}

func parseThrowTrue(line string) int {
	return parseSingleIntAtPosition(line, 29)
}

func parseThrowFalse(line string) int {
	return parseSingleIntAtPosition(line, 30)
}

func parseSingleIntAtPosition(line string, position int) int {
	value, _ := strconv.Atoi(line[position:])
	return int(value)
}
