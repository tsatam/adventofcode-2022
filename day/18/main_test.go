package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input string
		want  []Point3D
	}{
		{
			input: `1,1,1
2,1,1`,
			want: []Point3D{
				{1, 1, 1},
				{2, 1, 1},
			},
		},
		{
			input: `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`,
			want: []Point3D{
				{2, 2, 2},
				{1, 2, 2},
				{3, 2, 2},
				{2, 1, 2},
				{2, 3, 2},
				{2, 2, 1},
				{2, 2, 3},
				{2, 2, 4},
				{2, 2, 6},
				{1, 2, 5},
				{3, 2, 5},
				{2, 1, 5},
				{2, 3, 5},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := parseInput(tt.input)

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got [%v], want [%v]", got, tt.want)
			}
		})
	}
}

func TestCalcSurfaceArea(t *testing.T) {
	tests := []struct {
		cubes []Point3D
		want  int
	}{
		{
			cubes: []Point3D{
				{1, 1, 1},
				{2, 1, 1},
			},
			want: 10,
		},
		{
			cubes: []Point3D{
				{2, 2, 2},
				{1, 2, 2},
				{3, 2, 2},
				{2, 1, 2},
				{2, 3, 2},
				{2, 2, 1},
				{2, 2, 3},
				{2, 2, 4},
				{2, 2, 6},
				{1, 2, 5},
				{3, 2, 5},
				{2, 1, 5},
				{2, 3, 5},
			},
			want: 64,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := calcSurfaceArea(tt.cubes)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}
		})
	}
}

func TestCalcSurfaceAreaNoInterior(t *testing.T) {
	tests := []struct {
		cubes []Point3D
		want  int
	}{
		{
			cubes: []Point3D{
				{2, 2, 2},
				{1, 2, 2},
				{3, 2, 2},
				{2, 1, 2},
				{2, 3, 2},
				{2, 2, 1},
				{2, 2, 3},
				{2, 2, 4},
				{2, 2, 6},
				{1, 2, 5},
				{3, 2, 5},
				{2, 1, 5},
				{2, 3, 5},
			},
			want: 58,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := calcSurfaceAreaNoInterior(tt.cubes)

			if got != tt.want {
				t.Errorf("got [%d], want [%d]", got, tt.want)
			}
		})
	}
}
