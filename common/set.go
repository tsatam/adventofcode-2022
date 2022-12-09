package set

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
