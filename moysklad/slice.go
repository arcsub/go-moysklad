package moysklad

type Slice[E any] []*E

// S возвращает простой срез.
func (s Slice[E]) S() []*E {
	return s
}

// Iter возвращает итератор.
func (s Slice[E]) Iter() *Iterator[E] {
	return &Iterator[E]{el: s}
}

// Push добавляет элементы в конец среза.
func (s *Slice[E]) Push(elements ...*E) *Slice[E] {
	*s = append(*s, elements...)
	return s
}

// Unshift добавляет элементы в начало среза.
func (s *Slice[E]) Unshift(elements ...*E) *Slice[E] {
	*s = append(s.S()[:len(elements)], append(elements, s.S()[len(elements):]...)...)
	return s
}

func (s *Slice[E]) Pop() *E {
	sl := s.S()
	e := sl[len(*s)-1]
	*s = sl[:len(sl)-1]
	return e
}

func (s *Slice[E]) Shift() *E {
	sl := s.S()
	e := sl[0]
	*s = sl[1:]
	return e
}

// Filter фильтрация элементов.
func (s Slice[E]) Filter(f func(e *E) bool) Slice[E] {
	b := s[:0]
	for _, x := range s {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}
