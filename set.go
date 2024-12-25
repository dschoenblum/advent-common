package common

import "iter"

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] {
	s := &Set[T]{values: make(map[T]struct{})}
	s.Set(values...)
	return s
}

func (s *Set[T]) Set(values ...T) {
	for _, value := range values {
		s.values[value] = struct{}{}
	}
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.values[value]
	return ok
}

func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range s.values {
			if !yield(value) {
				break
			}
		}
	}
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Random() T {
	for value := range s.values {
		return value
	}
	var zero T
	return zero
}
