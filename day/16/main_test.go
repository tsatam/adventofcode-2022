package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`
	want := map[string]Valve{
		"AA": {name: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
		"BB": {name: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
		"CC": {name: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
		"DD": {name: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
		"EE": {name: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
		"FF": {name: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
		"GG": {name: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
		"HH": {name: "HH", flowRate: 22, tunnels: []string{"GG"}},
		"II": {name: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
		"JJ": {name: "JJ", flowRate: 21, tunnels: []string{"II"}},
	}

	got := parseInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func TestFindMaxPressure(t *testing.T) {

	valves = map[string]Valve{
		"AA": {name: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
		"BB": {name: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
		"CC": {name: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
		"DD": {name: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
		"EE": {name: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
		"FF": {name: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
		"GG": {name: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
		"HH": {name: "HH", flowRate: 22, tunnels: []string{"GG"}},
		"II": {name: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
		"JJ": {name: "JJ", flowRate: 21, tunnels: []string{"II"}},
	}
	want := 1651

	got := findMaxPressure()

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestFindMaxPressureWithElephant(t *testing.T) {

	valves = map[string]Valve{
		"AA": {name: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
		"BB": {name: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
		"CC": {name: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
		"DD": {name: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
		"EE": {name: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
		"FF": {name: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
		"GG": {name: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
		"HH": {name: "HH", flowRate: 22, tunnels: []string{"GG"}},
		"II": {name: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
		"JJ": {name: "JJ", flowRate: 21, tunnels: []string{"II"}},
	}
	want := 1707

	got := findMaxPressureWithElephant()

	if got != want {
		t.Errorf("got [%d], want [%d]", got, want)
	}
}

func TestFindAllShortestPaths(t *testing.T) {

	valves = map[string]Valve{
		"AA": {name: "AA", flowRate: 0, tunnels: []string{"DD", "II", "BB"}},
		"BB": {name: "BB", flowRate: 13, tunnels: []string{"CC", "AA"}},
		"CC": {name: "CC", flowRate: 2, tunnels: []string{"DD", "BB"}},
		"DD": {name: "DD", flowRate: 20, tunnels: []string{"CC", "AA", "EE"}},
		"EE": {name: "EE", flowRate: 3, tunnels: []string{"FF", "DD"}},
		"FF": {name: "FF", flowRate: 0, tunnels: []string{"EE", "GG"}},
		"GG": {name: "GG", flowRate: 0, tunnels: []string{"FF", "HH"}},
		"HH": {name: "HH", flowRate: 22, tunnels: []string{"GG"}},
		"II": {name: "II", flowRate: 0, tunnels: []string{"AA", "JJ"}},
		"JJ": {name: "JJ", flowRate: 21, tunnels: []string{"II"}},
	}

	want := map[string]map[string]int{
		"AA": {"BB": 1, "CC": 2, "DD": 1, "EE": 2, "HH": 5, "JJ": 2},
		"BB": {"BB": 0, "CC": 1, "DD": 2, "EE": 3, "HH": 6, "JJ": 3},
		"CC": {"BB": 1, "CC": 0, "DD": 1, "EE": 2, "HH": 5, "JJ": 4},
		"DD": {"BB": 2, "CC": 1, "DD": 0, "EE": 1, "HH": 4, "JJ": 3},
		"EE": {"BB": 3, "CC": 2, "DD": 1, "EE": 0, "HH": 3, "JJ": 4},
		"HH": {"BB": 6, "CC": 5, "DD": 4, "EE": 3, "HH": 0, "JJ": 7},
		"JJ": {"BB": 3, "CC": 4, "DD": 3, "EE": 4, "HH": 7, "JJ": 0},
	}

	got := findAllShortestPaths("AA")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got [%v], want [%v]", got, want)
	}

}
