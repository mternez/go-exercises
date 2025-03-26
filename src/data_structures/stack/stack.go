package stack

type Element[T any] struct {
	value    T
	next     *Element[T]
	previous *Element[T]
}

func (e *Element[T]) Value() T {
	return e.value
}

type Stack[T any] struct {
	top *Element[T]
}

func (s *Stack[E]) Push(element E) {
	if s.top != nil {
		s.top.next = &Element[E]{value: element, previous: s.top}
		s.top = s.top.next
	} else {
		s.top = &Element[E]{value: element}
	}
}

func (s *Stack[E]) Pop() *E {
	if s.top != nil {
		previousHead := s.top
		if previousHead.previous != nil {
			s.top = previousHead.previous
			s.top.next = nil
		} else {
			s.top = nil
		}
		previousHead.next = nil
		previousHead.previous = nil
		return &previousHead.value
	}
	return nil
}

func (s *Stack[E]) Peek() E {
	return s.top.value
}

func (s *Stack[E]) Empty() bool {
	return s.top == nil
}
