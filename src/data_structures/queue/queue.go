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
	head *Element[T]
	tail *Element[T]
}

func (s *Queue[E]) Push(element E) {

	e := &Element[E]{value: element}
	if s.head == nil {
		s.head = e
		s.tail = e
	} else {
		s.tail.next = e
		e.previous = s.tail
		s.tail = e
	}
}

func (s *Queue[E]) Pop() *E {

	if s.head != nil {

		previousHead := s.head
		previousTail := s.tail

		s.head = s.tail

		if previousTail.previous != nil {
			s.tail = previousTail.previous
			s.tail.next = nil
			previousTail.previous = nil
		} else {
			s.tail = nil
			s.head = nil
		}

		return &previousHead.value
	}

	return nil
}

func (s *Queue[E]) Peek() E {
	var value E
	if s.tail != nil {
		value = s.tail.value
	}
	return value
}

func (s *Queue[E]) Empty() bool {
	return s.head == nil
}

func (s *Queue[E]) Size() int {
	size := 0
	tail := s.tail
	if tail != nil {
		size++
	}
	for tail.previous != nil {
		tail = tail.previous
		if tail != nil {
			size++
		}
	}
	return size
}
