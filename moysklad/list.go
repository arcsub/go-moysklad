package moysklad

type List[T any] struct {
	Context Context        `json:"context"`
	Meta    MetaCollection `json:"meta"`
	Rows    Slice[T]       `json:"rows"`
}

func (l List[T]) String() string {
	return Stringify(l)
}
