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
		if item, err := strconv.Atoi(rawItem); err == nil {
			items[i] = item
		} else {
			log.Fatal(err)
		}
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
		return AddOperation{value}
	}
	if rawOperation[4] == '*' {
		return MulOperation{value}
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
	if value, err := strconv.Atoi(line[position:]); err == nil {
		return value
	} else {
		log.Fatalf("Error parsing line [%s], position [%d], err [%v]\n", line, position, err)
		return -1
	}
}
