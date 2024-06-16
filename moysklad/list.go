package moysklad

type List[T any] struct {
	Context Context `json:"context"`
	MetaArray[T]
}

func (list List[T]) String() string {
	return Stringify(list)
}
