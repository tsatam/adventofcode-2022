package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("pt1: %v\n", findFirstNUnique(input, 4))
	fmt.Printf("pt2: %v\n", findFirstNUnique(input, 14))
}

func findFirstNUnique(input string, wantDistinct int) int {
	r := []rune(input)

	for i := wantDistinct; i <= len(r); i++ {
		if s := NewSet(r[i-wantDistinct : i]...); s.Size() == wantDistinct {
			return i
		}
	}
	return -1
}

type set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items ...T) set[T] {
	s := set[T]{make(map[T]struct{})}

	s.AddAll(items...)

	return s
}

func (s *set[T]) Add(item T) {
	s.m[item] = struct{}{}
}

func (s *set[T]) AddAll(items ...T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *set[T]) Size() int {
	return len(s.m)
}
