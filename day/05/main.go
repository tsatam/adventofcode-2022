package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input              string
	instructionMatcher = regexp.MustCompile(`move (?P<amount>\d+) from (?P<from>\d) to (?P<to>\d)`)
)

func main() {
	ship := parseInput(input)
	ship.processInstructions9000()
	fmt.Printf("pt1: %v\n", ship.outputHighestCrates())
	ship = parseInput(input)
	ship.processInstructions9001()
	fmt.Printf("pt2: %v\n", ship.outputHighestCrates())
}

func parseInput(input string) Ship {
	splitInput := strings.Split(input, "\n\n")

	return Ship{
		crates:       parseCrates(splitInput[0]),
		instructions: parseInstructions(splitInput[1]),
	}
}

func parseCrates(raw string) [][]byte {
	lines := strings.Split(raw, "\n")

	crates := make([][]byte, getNumCrates(lines[len(lines)-1]))
	for i := range crates {
		crates[i] = make([]byte, 0)
	}

	for lineIndex := len(lines) - 2; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]
		for i := 0; i < len(crates); i++ {
			rawCrate := line[i*4 : i*4+3]
			if rawCrate != "   " {
				crate := parseCrate(rawCrate)
				crates[i] = append(crates[i], crate)
			}
		}
	}

	return crates
}

func parseInstructions(raw string) []Instruction {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	return instructions
}

func parseInstruction(raw string) Instruction {
	match := instructionMatcher.FindStringSubmatch(raw)
	values := make([]int, 3)

	for i := 1; i < 4; i++ {
		parsed, err := strconv.Atoi(match[i])
		if err != nil {
			log.Fatal(err)
		}
		values[i-1] = parsed
	}

	return Instruction{
		amount: values[0],
		from:   values[1],
		to:     values[2],
	}
}

func parseCrate(raw string) byte {
	return raw[1]
}

func getNumCrates(lastLine string) int {
	split := strings.Split(strings.TrimSpace(lastLine), "   ")
	lastRaw := split[len(split)-1]
	last, err := strconv.Atoi(lastRaw)
	if err != nil {

		log.Fatalf("Error parsing num crates: [%v], lastLine: [%v], split: [%v], lastRaw: [%v]", err, lastLine, split, lastRaw)
	}
	return last
}

type Ship struct {
	crates       [][]byte
	instructions []Instruction
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func (s *Ship) processInstructions9000() {
	for _, instruction := range s.instructions {
		s.processInstruction9000(instruction)
	}
}
func (s *Ship) processInstructions9001() {
	for _, instruction := range s.instructions {
		s.processInstruction9001(instruction)
	}
}

func (s *Ship) processInstruction9000(instruction Instruction) {
	for count := 0; count < instruction.amount; count++ {
		fromIdx := instruction.from - 1
		from := s.crates[fromIdx]
		toIdx := instruction.to - 1

		crate := from[len(from)-1]
		s.crates[fromIdx] = from[:len(from)-1]
		s.crates[toIdx] = append(s.crates[toIdx], crate)
	}
}

func (s *Ship) processInstruction9001(instruction Instruction) {
	fromIdx := instruction.from - 1
	from := s.crates[fromIdx]
	toIdx := instruction.to - 1

	cratesToMove := from[len(from)-instruction.amount:]
	s.crates[fromIdx] = from[:len(from)-instruction.amount]
	s.crates[toIdx] = append(s.crates[toIdx], cratesToMove...)
}

func (s *Ship) outputHighestCrates() string {
	highestCrates := make([]byte, len(s.crates))
	for i, stack := range s.crates {
		highestCrate := stack[len(stack)-1]
		highestCrates[i] = highestCrate
	}
	return string(highestCrates)
}
