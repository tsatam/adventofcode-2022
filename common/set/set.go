package set

import "fmt"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{make(map[T]struct{})}

	s.AddAll(items...)

	return s
}

func (s *Set[T]) Add(item T) {
	s.m[item] = struct{}{}
}

func (s *Set[T]) AddAll(items ...T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set[T]) GetOne() (T, error) {
	for k := range s.m {
		return k, nil
	}
	var empty T
	return empty, fmt.Errorf("Set [%v] empty", s)
}

func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}

func (s *Set[T]) Slice() []T {
	keys := make([]T, len(s.m))
	i := 0
	for k, _ := range s.m {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
