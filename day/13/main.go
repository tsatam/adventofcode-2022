package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input string
)

func main() {
	sum := solve(input)
	fmt.Printf("pt1: [%d]\n", sum)
}

func solve(input string) int {
	sum := 0

	pairs := strings.Split(input, "\n\n")

	for i, pair := range pairs {
		splitPair := strings.Split(pair, "\n")
		left := splitPair[0]
		right := splitPair[1]

		if compareSignals(left, right) {
			sum += i + 1
		}
	}

	return sum
}

func compareSignals(left, right string) bool {
	return compareSignalsInternal(left, right) <= 0
}

func compareSignalsInternal(left, right string) int {
	if isInt(left) && isInt(right) {
		li, _ := strconv.Atoi(left)
		ri, _ := strconv.Atoi(right)

		if li < ri {
			return -1
		} else if li == ri {
			return 0
		} else {
			return 1
		}
	} else {
		ll := parseList(left)
		rl := parseList(right)

		for i := 0; i < min(len(ll), len(rl)); i++ {
			result := compareSignalsInternal(ll[i], rl[i])
			if result != 0 {
				return result
			}
		}

		if len(ll) < len(rl) {
			return -1
		} else if len(rl) < len(ll) {
			return +1
		} else {
			return 0
		}
	}
}

func isInt(s string) bool {
	return !strings.ContainsAny(s, "[],")
}

func parseList(s string) []string {
	if isInt(s) {
		return []string{s}
	}

	res := make([]string, 0)

	var sb strings.Builder

	parenStack := 0
	for _, c := range s[1 : len(s)-1] {
		if c == ',' && parenStack == 0 {
			res = append(res, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(c)
			if c == '[' {
				parenStack++
			} else if c == ']' {
				parenStack--
			}
		}
	}
	str := sb.String()
	if str != "" {
		res = append(res, str)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
