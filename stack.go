package common

type Stack[T any] []T

func NewStack[T any](items ...T) *Stack[T] {
	s := &Stack[T]{}
	*s = append(*s, items...)
	return s
}

func (s *Stack[T]) Push(items ...T) {
	*s = append(*s, items...)
}

func (s *Stack[T]) Top() (item T) {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) TryTop() (item T, ok bool) {
	if !s.IsEmpty() {
		item = (*s)[len(*s)-1]
		ok = true
	}
	return
}

func (s *Stack[T]) Pop() (item T) {
	item = s.Top()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack[T]) TryPop() (item T, ok bool) {
	top, ok := s.TryTop()
	if ok {
		*s = (*s)[:len(*s)-1]
	}
	return top, ok
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Count() int {
	return len(*s)
}
