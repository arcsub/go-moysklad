package moysklad

type List[T any] struct {
	Context Context `json:"context"`
	MetaArray[T]
}
