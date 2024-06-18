package moysklad

type Slice[E any] []*E

func NewSlice[T any]() Slice[T] {
	return make(Slice[T], 0)
}

func NewSliceFrom[T any](elements []T) Slice[T] {
	s := make(Slice[T], len(elements))
	for _, element := range elements {
		s.Push(&element)
	}
	return s
}

// S возвращает простой срез.
func (slice Slice[E]) S() []*E {
	return slice
}

// Len возвращает количество элементов
func (slice Slice[E]) Len() int {
	return len(slice)
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
	for chunkSize < items.Len() {
		items, chunks = items[chunkSize:], append(chunks, items[:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func (slice Slice[E]) UnPtr() []E {
	var e []E
	for _, row := range slice {
		e = append(e, Deref(row))
	}
	return e
}

// AsMetaWrapper приводит элементы слайса к типу MetaWrapper
func (slice Slice[E]) AsMetaWrapper() []MetaWrapper {
	var mw []MetaWrapper
	for _, elem := range slice {
		if m, ok := any(elem).(MetaOwner); ok {
			mw = append(mw, m.GetMeta().Wrap())
		}
	}
	return mw
}
