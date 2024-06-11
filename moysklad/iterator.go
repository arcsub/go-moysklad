package moysklad

import (
	"encoding/json"
	"sync"
)

// Iterator структура итератора
type Iterator[E any] struct {
	el  Slice[E]
	idx int
	mu  sync.Mutex
}

// Len возвращает количество элементов
func (iterator *Iterator[E]) Len() int {
	return len(iterator.el)
}

// HasNext возвращает true, если текущее значение счётчика меньше длины среза
func (iterator *Iterator[E]) HasNext() bool {
	return iterator.idx < iterator.Len()
}

// Next возвращает следующий элемент среза или nil, если элемент отсутствует
func (iterator *Iterator[E]) Next() *E {
	if iterator.HasNext() {
		iterator.mu.Lock()
		defer iterator.mu.Unlock()
		row := iterator.el[iterator.idx]
		iterator.idx += 1
		return row
	}
	return nil
}

// Push добавляет элементы в срез
func (iterator *Iterator[E]) Push(elements ...*E) {
	iterator.mu.Lock()
	defer iterator.mu.Unlock()
	iterator.el = append(iterator.el, elements...)
}

// Stop сбрасывает текущее значение индекса
func (iterator *Iterator[E]) Stop() {
	iterator.idx = 0
}

// Filter фильтрация элементов итератора.
// Возвращает новый итератор с отфильтрованными элементами.
func (iterator *Iterator[E]) Filter(f func(e *E) bool) *Iterator[E] {
	var n = &Iterator[E]{}

	iterator.mu.Lock()
	defer iterator.mu.Unlock()

	for i := 0; i < iterator.Len(); i++ {
		el := iterator.el[i]
		if f(el) {
			n.Push(el)
		}
	}

	return n
}

// Slice возвращает срез элементов
func (iterator *Iterator[E]) Slice() Slice[E] {
	return iterator.el
}

// MarshalJSON implements the json.Marshaler interface.
func (iterator *Iterator[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(iterator.el)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (iterator *Iterator[E]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &iterator.el)
}
