package moysklad

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Iterator структура итератора
type Iterator[E any] struct {
	idx  int // счётчик
	size int // максимальный лимит среза (при 0 лимит не установлен)
	mu   sync.Mutex
	el   Slice[E]
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
		return &row
	}
	return nil
}

// Push добавляет элементы в срез
func (r *Iterator[E]) Push(elements ...E) error {
	if r.size > 0 {
		if len(r.el) > r.size {
			return fmt.Errorf("количество добавляемых элементов превышает допустимый лимит: %d", r.size)
		}

		limit := r.size - len(r.el)
		if limit < 0 {
			limit = 0
		}
		elements = elements[:limit]
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.el = append(r.el, elements...)
	return nil
}

// PushPtr добавляет элементы в срез
func (r *Iterator[E]) PushPtr(elements ...*E) error {
	var elements2 []E
	for _, element := range elements {
		elements2 = append(elements2, *element)
	}
	return r.Push(elements2...)
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

type Slice[E any] []E

// Iter возвращает итератор
func (s Slice[E]) Iter(opt ...int) *Iterator[E] {
	var size int
	if len(opt) > 0 {
		size = opt[0]
	}
	return &Iterator[E]{el: s, size: size}
}
