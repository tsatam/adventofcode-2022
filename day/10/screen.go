package main

import (
	"bytes"
	"common/cartesian"
	"log"
	"reflect"
)

const (
	charWidth = 5
)

func iterToScreenPosition(i int) cartesian.Point {
	x := i % 40
	y := i / 40

	return cartesian.Point{X: x, Y: y}
}

type Screen = [6][40]uint8
type ScreenChar = [6][5]uint8

var screenChars = map[rune]ScreenChar{
	'A': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
	},
	'B': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	},
	'C': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'E': {
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	},
	'F': {
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'G': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
	},
	'H': {
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
	},
	'I': {
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0},
	},
	'J': {
		{0, 0, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'K': {
		{1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 0, 1, 0},
	},
	'L': {
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	},
	'O': {
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'P': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
	},
	'R': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 0, 1, 0},
	},
	'S': {
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	},
	'U': {
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
	},
	'Y': {
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 1, 0, 0},
	},
	'Z': {
		{1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	},
}

func charAt(screen Screen, pos int) rune {
	screenChar := getScreenChar(screen, pos)

	for char, screenCharToMatch := range screenChars {
		if reflect.DeepEqual(screenChar, screenCharToMatch) {
			return char
		}
	}
	log.Fatal("Could not find char for screenChar [%v]", screenChar)
	return '_'
}

func getScreenChar(screen Screen, pos int) ScreenChar {
	var screenChar ScreenChar
	for y := 0; y < 6; y++ {
		for x := 0; x < 5; x++ {
			screenChar[y][x] = screen[y][pos*5+x]
		}
	}
	return screenChar
}

func print(s Screen) string {
	var buffer bytes.Buffer

	for i := 0; i < 8; i++ {
		buffer.WriteRune(charAt(s, i))
	}

	result := buffer.String()
	return result
}

func debug(s Screen) string {
	var buffer bytes.Buffer

	lit := '#'
	dark := '.'

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == 1 {
				buffer.WriteRune(lit)
			} else {
				buffer.WriteRune(dark)
			}
		}
		buffer.WriteRune('\n')
	}
	result := buffer.String()
	return result
}
