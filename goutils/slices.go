package goutils

type Iterator[T any] struct {
	fields   []T
	length   int
	position int
	called   bool
}

func NewIterator[T any](s []T) Iterator[T] {
	return Iterator[T]{
		fields: s,
		length: len(s),
	}
}

func (i *Iterator[T]) FieldsLeft() []T {
	return i.fields[i.position:]
}

func (i *Iterator[T]) LenLeft() int {
	return i.length - 1 - i.position
}

func (i *Iterator[T]) Length() int {
	return i.length
}

func (i *Iterator[T]) Position() int {
	return i.position
}

func (i *Iterator[T]) Reset() {
	i.position = 0
	i.called = false
}

func (i *Iterator[T]) Next() bool {
	if i.position >= i.length-1 {
		return false
	}
	if !i.called {
		i.called = true
		if i.position == 0 {
			return true
		}
	}
	i.position++
	return true
}

func (i *Iterator[T]) Prev() bool {
	if i.position > 0 {
		i.position--
		return true
	}
	return false
}

func (i *Iterator[T]) Skip(n int) bool {
	i.position += n
	if i.position < 0 {
		i.Reset()
		return false
	} else if i.position >= i.length {
		i.position = i.length - 1
		return false
	}
	return true
}

func (i *Iterator[T]) Value() T {
	if !i.called {
		i.called = true
	}
	return i.fields[i.position]
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	v, h := s.Top()
	*s = (*s)[:h]
	return v
}

func (s *Stack[T]) Top() (T, int) {
	h := len(*s) - 1
	if h < 0 {
		return *new(T), 0
	}
	return (*s)[h], h
}
