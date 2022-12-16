package main

import (
	maxpq "common/max_priority_queue"
	minpq "common/min_priority_queue"
	"common/set"
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	Time = 30
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
	startTimeRemaining := 30
	startValveState := make(ValveState, len(valves))

	shortestPaths := findAllShortestPaths(startPosition)

	dests := set.NewSet[string]()
	for dest := range shortestPaths[startPosition] {
		dests.Add(dest)
		startValveState[dest] = -1
	}

	return findMaxPressureInternal(startPosition, startTimeRemaining, startValveState, dests, shortestPaths)
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
	startState ValveState,
	destinations set.Set[string],
	shortestPaths map[string]map[string]int,
) int {

	if startTimeRemaining == 0 {
		return 0
	}

	queue := maxpq.New[Action](0, len(valves))

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
		if action.timeRemaining == 0 || allImportantValvesOpen(currState, destinations) {
			continue
		}

		shouldTraverse := make([]string, len(shortestPaths[action.valve]))
		for neighbor := range shortestPaths[action.valve] {
			if neighbor != action.valve {
				shouldTraverse = append(shouldTraverse, neighbor)
			}
		}
		sort.Slice(shouldTraverse, func(i, j int) bool {
			return shortestPaths[action.valve][shouldTraverse[i]] < shortestPaths[action.valve][shouldTraverse[j]]
		})

		for _, neighbor := range shouldTraverse {
			distance := shortestPaths[action.valve][neighbor]

			if currState[neighbor] == -1 && distance <= action.timeRemaining {
				reachedAt := action.timeRemaining - distance - 1
				newState := openValve(currState, neighbor, reachedAt)
				newPressure := calcPressure(newState)

				newAction := Action{
					valve:         neighbor,
					state:         &newState,
					timeRemaining: reachedAt,
					pressure:      newPressure,
				}
				queue.AddAtPriority(newAction, action.pressure)
			}
		}
	}

	return maxPressure
}

func findShortestPaths(source string) map[string]int {
	queue := minpq.New[string](0, len(valves))
	shortestDistances := make(map[string]int, len(valves))

	for k := range valves {
		if k == source {
			shortestDistances[k] = 0
		} else {
			shortestDistances[k] = len(valves)
		}

		queue.AddAtPriority(k, shortestDistances[k])
	}

	for !queue.Empty() {
		next := queue.PopMin()

		currentDistance := shortestDistances[next] + 1

		for _, neighbor := range valves[next].tunnels {
			if currentDistance < shortestDistances[neighbor] {
				shortestDistances[neighbor] = currentDistance
				queue.SetPriority(neighbor, currentDistance)
			}
		}
	}

	return shortestDistances
}

func calcPressure(state ValveState) int {
	pressure := 0
	for valve, openedAt := range state {
		if openedAt > 0 {
			pressure += valves[valve].flowRate * openedAt
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

func openValve(state ValveState, position string, time int) ValveState {
	newState := make(ValveState, len(state))
	for k, v := range state {
		if k == position {
			newState[k] = time
		} else {
			newState[k] = v
		}
	}
	return newState
}
