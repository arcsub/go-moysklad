package moysklad

import (
	"encoding/json"
	"sync"
)

// Iterator структура итератора
type Iterator[E any] struct {
	idx int // счётчик
	mu  sync.Mutex
	el  Slice[E]
}

// Len возвращает количество элементов
func (r Iterator[E]) Len() int {
	return len(r.el)
}

// HasNext возвращает true, если текущее значение счётчика меньше длины среза
func (r Iterator[E]) HasNext() bool {
	return r.idx < r.Len()
}

// Next возвращает следующий элемент среза или nil, если элемент отсутствует
func (r *Iterator[E]) Next() *E {
	if r.HasNext() {
		r.mu.Lock()
		defer r.mu.Unlock()
		row := r.el[r.idx]
		r.idx += 1
		return row
	}
	return nil
}

// Push добавляет элементы в срез
func (r *Iterator[E]) Push(elements ...*E) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.el = append(r.el, elements...)
	return nil
}

// Stop сбрасывает текущее значение индекса
func (r *Iterator[E]) Stop() {
	r.idx = 0
}

// Slice возвращает срез элементов
func (r *Iterator[E]) Slice() Slice[E] {
	return r.el
}

func (r Iterator[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.el)
}

func (r *Iterator[E]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.el)
}

type Slice[E any] []*E

// Iter возвращает итератор
func (s Slice[E]) Iter() *Iterator[E] {
	return &Iterator[E]{el: s}
}

// Push добавляет элементы в срез.
func (s *Slice[E]) Push(elements ...*E) *Slice[E] {
	*s = append(*s, elements...)
	return s
}
