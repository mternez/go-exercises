package queue

type Element[T any] struct {
	value    T
	next     *Element[T]
	previous *Element[T]
}

func (e *Element[T]) Value() T {
	return e.value
}

type Queue[T any] struct {
	top    *Element[T]
	bottom *Element[T]
}

func (s *Queue[E]) Push(element E) {

	e := &Element[E]{value: element}
	if s.top == nil {
		s.top = e
		s.bottom = e
	} else {
		s.bottom.next = e
		e.previous = s.bottom
		s.bottom = e
	}
}

func (s *Queue[E]) Pop() *E {

	if s.top != nil {

		previousTop := s.top

		previousBottom := s.bottom

		s.top = s.bottom

		if previousBottom.previous != nil {
			s.bottom = previousBottom.previous
			s.bottom.next = nil
			previousBottom.previous = nil
		}

		return &previousTop.value
	}
	return nil
}

func (s *Queue[E]) Peek() E {
	return s.bottom.value
}

func (s *Queue[E]) Empty() bool {
	return s.top == nil
}

func (s *Queue[E]) Size() int {
	size := 0
	bottom := s.bottom
	if bottom != nil {
		size++
	}
	for bottom.previous != nil {
		bottom = bottom.previous
		if bottom != nil {
			size++
		}
	}
	return size
}
