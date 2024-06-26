package moysklad

// Slice представляет собой срез на указатели типа T.
type Slice[E any] []*E

// NewSlice возвращает новый [Slice] с нулевой длиной.
func NewSlice[T any]() Slice[T] {
	return make(Slice[T], 0)
}

// NewSliceFrom конвертирует слайс типов T в тип [Slice].
func NewSliceFrom[T any](elements []T) Slice[T] {
	s := make(Slice[T], 0, len(elements))
	for _, element := range elements {
		s.Push(&element)
	}
	return s
}

// S преобразует обратно [Slice] в слайс указателей на тип T и возвращает его.
func (slice Slice[E]) S() []*E {
	return slice
}

// Len возвращает количество элементов.
func (slice Slice[E]) Len() int {
	return len(slice)
}

// Iter возвращает итератор [Iterator].
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

// Pop удаляет последний элемент из слайса и возвращает удалённое значение.
func (slice *Slice[E]) Pop() *E {
	sl := slice.S()
	e := sl[len(*slice)-1]
	*slice = sl[:len(sl)-1]
	return e
}

// Shift удаляет элемент по нулевому индексу, сдвигает значения по последовательным индексам вниз,
// затем возвращает удалённое значение.
func (slice *Slice[E]) Shift() *E {
	sl := slice.S()
	e := sl[0]
	*slice = sl[1:]
	return e
}

// Filter принимает функцию типа func(e *E) bool и применяет её в качестве фильтрации элементов.
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

// UnPtr преобразует обратно [Slice] в слайс разыменованных типов T и возвращает его.
func (slice Slice[E]) UnPtr() []E {
	var e = make([]E, 0, slice.Len())
	for _, row := range slice {
		e = append(e, Deref(row))
	}
	return e
}

// AsMetaWrapper преобразует элементы слайса к слайсу типов [MetaWrapper].
func (slice Slice[E]) AsMetaWrapper() []MetaWrapper {
	var mw = make([]MetaWrapper, 0, slice.Len())
	for _, elem := range slice {
		if m, ok := any(elem).(MetaOwner); ok {
			mw = append(mw, m.GetMeta().Wrap())
		}
	}
	return mw
}
