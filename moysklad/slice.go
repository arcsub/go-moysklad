package moysklad

type Slice[E any] []*E

// S возвращает простой срез.
func (slice Slice[E]) S() []*E {
	return slice
}

// Iter возвращает итератор.
func (slice Slice[E]) Iter() *Iterator[E] {
	return &Iterator[E]{el: slice}
}

// Push добавляет элементы в конец среза.
func (slice *Slice[E]) Push(elements ...*E) *Slice[E] {
	*slice = append(*slice, elements...)
	return slice
}

// Unshift добавляет элементы в начало среза.
func (slice *Slice[E]) Unshift(elements ...*E) *Slice[E] {
	*slice = append(slice.S()[:len(elements)], append(elements, slice.S()[len(elements):]...)...)
	return slice
}

func (slice *Slice[E]) Pop() *E {
	sl := slice.S()
	e := sl[len(*slice)-1]
	*slice = sl[:len(sl)-1]
	return e
}

func (slice *Slice[E]) Shift() *E {
	sl := slice.S()
	e := sl[0]
	*slice = sl[1:]
	return e
}

// Filter фильтрация элементов.
func (slice Slice[E]) Filter(f func(e *E) bool) Slice[E] {
	b := slice[:0]
	for _, x := range slice {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

// IntoChunks разбивает слайс на разные слайсы с заданной размерностью.
func (slice Slice[E]) IntoChunks(chunkSize int) (chunks []Slice[E]) {
	var items = slice
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
