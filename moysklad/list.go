package moysklad

type List[T any] struct {
	Context Context        `json:"context"`
	Rows    Slice[T]       `json:"rows"`
	Meta    MetaCollection `json:"meta"`
}

func (l List[T]) String() string {
	return Stringify(l)
}
