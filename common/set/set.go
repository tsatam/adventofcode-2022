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

func (s *Set[T]) Union(other Set[T]) {
	for item := range other.m {
		s.Add(item)
	}
}

func (s *Set[T]) Subtraction(other Set[T]) {
	for item := range other.m {
		s.Remove(item)
	}
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set[T]) ContainsAll(other Set[T]) bool {
	for item := range other.m {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s *Set[T]) GetOne() (T, error) {
	for k := range s.m {
		return k, nil
	}
	var empty T
	return empty, fmt.Errorf("Set [%v] empty", s)
}

func (s *Set[T]) PopOne() (T, error) {
	one, err := s.GetOne()
	if err != nil {
		var empty T
		return empty, err
	}
	s.Remove(one)
	return one, nil
}

func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}

func (s *Set[T]) Slice() []T {
	keys := make([]T, len(s.m))
	i := 0
	for k := range s.m {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

func (s *Set[T]) Copy() Set[T] {
	new := Set[T]{make(map[T]struct{}, len(s.m))}

	for k := range s.m {
		new.m[k] = struct{}{}
	}

	return new
}
