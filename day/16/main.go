package main

import (
	maxpq "common/max_priority_queue"
	minpq "common/min_priority_queue"
	"common/set"
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed input
	input       string
	lineMatcher = regexp.MustCompile(`Valve (?P<name>..) has flow rate=(?P<flowRate>\d+); tunnels? leads? to valves? (?P<tunnels>(?:.., )*..).*`)
	valves      map[string]Valve
)

func main() {
	parseInput(input)

	maxPressure := findMaxPressure()
	fmt.Printf("pt1: [%d]\n", maxPressure)

	maxPressureWithElephant := findMaxPressureWithElephant()
	fmt.Printf("pt2: [%d]\n", maxPressureWithElephant)
}

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

func parseInput(input string) map[string]Valve {
	match := lineMatcher.FindAllStringSubmatch(input, -1)
	valves = make(map[string]Valve, len(match))

	for _, valve := range match {
		name := valve[lineMatcher.SubexpIndex("name")]
		flowRate, err := strconv.Atoi(valve[lineMatcher.SubexpIndex("flowRate")])
		if err != nil {
			log.Fatal(err)
		}
		tunnels := strings.Split(valve[lineMatcher.SubexpIndex("tunnels")], ", ")

		valves[name] = Valve{
			name:     name,
			flowRate: flowRate,
			tunnels:  tunnels,
		}
	}

	return valves
}

func findMaxPressure() int {

	startPosition := "AA"
	startTimeRemaining := int(30)

	shortestPaths := findAllShortestPaths(startPosition)

	dests := set.NewSet[string]()
	for dest := range shortestPaths[startPosition] {
		dests.Add(dest)
	}

	return findMaxPressureInternal(startPosition, startTimeRemaining, dests, shortestPaths)
}

func findMaxPressureWithElephant() int {
	startPosition := "AA"
	startTimeRemaining := 26

	shortestPaths := findAllShortestPaths(startPosition)

	dests := set.NewSet[string]()
	for dest := range shortestPaths[startPosition] {
		dests.Add(dest)
	}

	maxPressure := 0

	for lenSplit := 1; lenSplit < dests.Size()/2; lenSplit++ {
		permutations := getPermutations(dests, lenSplit)

		for _, permutation := range permutations {

			humanDests := set.NewSet(permutation...)
			elephantDests := dests.Copy()
			elephantDests.Subtraction(humanDests)

			humanPressure := findMaxPressureInternal(startPosition, startTimeRemaining, humanDests, shortestPaths)
			elephantPressure := findMaxPressureInternal(startPosition, startTimeRemaining, elephantDests, shortestPaths)

			pressure := humanPressure + elephantPressure

			if pressure > maxPressure {
				maxPressure = pressure
			}
		}
	}

	return maxPressure
}

func findAllShortestPaths(startPosition string) map[string]map[string]int {
	validStarts := set.NewSet(startPosition)
	validDests := set.NewSet[string]()

	for k, v := range valves {
		if v.flowRate > 0 {
			validStarts.Add(k)
			validDests.Add(k)
		}
	}

	shortestPaths := make(map[string]map[string]int, 0)
	for _, start := range validStarts.Slice() {
		shortestPaths[start] = make(map[string]int, 0)
		for dest, cost := range findShortestPaths(start) {
			if validDests.Contains(dest) {
				shortestPaths[start][dest] = cost
			}
		}
	}

	return shortestPaths
}

type Action struct {
	valve         string
	state         *ValveState
	pressure      int
	timeRemaining int
}

// stores when valve is opened. -1 indicates not open
type ValveState = map[string]int

func findMaxPressureInternal(
	startPosition string,
	startTimeRemaining int,
	destinations set.Set[string],
	shortestPaths map[string]map[string]int,
) int {

	destArr := destinations.Slice()
	startState := make(ValveState, destinations.Size())
	for _, dest := range destArr {
		startState[dest] = -1
	}

	queue := maxpq.New[Action](0, 1000000)

	queue.AddAtPriority(
		Action{
			valve:         startPosition,
			state:         &startState,
			timeRemaining: startTimeRemaining,
			pressure:      0,
		},
		0,
	)

	maxPressure := 0

	for !queue.Empty() {
		action := queue.PopMax()

		if action.pressure > maxPressure {
			maxPressure = action.pressure
		}

		currState := *action.state

		// abandon path if it is impossible to beat the current max pressure
		potentialPressure := calcPressureCeiling(*action.state, action.timeRemaining)
		if potentialPressure < maxPressure {
			continue
		}

		if action.timeRemaining == 0 || allImportantValvesOpen(currState, destinations) {
			continue
		}

		shouldTraverse := make([]string, 0)

		for _, neighbor := range destArr {
			if neighbor != action.valve && currState[neighbor] == -1 {
				shouldTraverse = append(shouldTraverse, neighbor)
			}
		}

		for _, neighbor := range shouldTraverse {
			distance := shortestPaths[action.valve][neighbor]

			if distance <= action.timeRemaining {
				reachedAt := action.timeRemaining - distance - 1
				newState := openValve(currState, neighbor, reachedAt)
				newPressure := calcPressure(newState)

				newAction := Action{
					valve:         neighbor,
					state:         &newState,
					timeRemaining: reachedAt,
					pressure:      newPressure,
				}
				queue.AddAtPriority(newAction, newPressure)
			}
		}
	}

	return maxPressure
}

func getPermutations[T comparable](s set.Set[T], len int) [][]T {
	if len == 0 {
		return [][]T{}
	}

	result := make([][]T, 0)

	for _, item := range s.Slice() {
		start := []T{item}

		if len == 1 {
			result = append(result, start)
		} else {
			copySet := s.Copy()
			copySet.Remove(item)

			for _, rest := range getPermutations(copySet, len-1) {
				next := append(start, rest...)
				result = append(result, next)
			}
		}
	}

	return result
}

func findShortestPaths(source string) map[string]int {
	queue := minpq.New[string](0, len(valves))
	shortestDistances := make(map[string]int, len(valves))

	for k := range valves {
		if k == source {
			shortestDistances[k] = 0
		} else {
			shortestDistances[k] = int(len(valves))
		}

		queue.AddAtPriority(k, int(shortestDistances[k]))
	}

	for !queue.Empty() {
		next := queue.PopMin()

		currentDistance := shortestDistances[next] + 1

		for _, neighbor := range valves[next].tunnels {
			if currentDistance < shortestDistances[neighbor] {
				shortestDistances[neighbor] = currentDistance
				queue.SetPriority(neighbor, int(currentDistance))
			}
		}
	}

	return shortestDistances
}

func calcPressure(state ValveState) int {
	pressure := 0
	for valve, openedAt := range state {
		if openedAt > 0 {
			pressure += valves[valve].flowRate * int(openedAt)
		}
	}

	return pressure
}

func calcPressureCeiling(state ValveState, timeRemaining int) int {
	pressure := 0
	for valve, openedAt := range state {
		if openedAt == -1 {
			pressure += valves[valve].flowRate * (int(timeRemaining) - 1)
		} else {
			pressure += valves[valve].flowRate * int(openedAt)
		}
	}
	return pressure
}

func allImportantValvesOpen(state ValveState, destinations set.Set[string]) bool {
	for _, destination := range destinations.Slice() {
		if state[destination] == -1 {
			return false
		}
	}

	return true
}

func openValve(state ValveState, position string, timeRemaining int) ValveState {
	// already opened
	if state[position] != -1 {
		return state
	}

	newState := make(ValveState, len(state))
	for k, v := range state {
		if k == position {
			newState[k] = int(timeRemaining)
		} else {
			newState[k] = v
		}
	}
	return newState
}
