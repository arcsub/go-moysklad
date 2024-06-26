package moysklad

// Context объект, содержащий метаданные о выполнившем запрос сотруднике.
type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}

// String реализует интерфейс [fmt.Stringer].
func (context Context) String() string {
	return Stringify(context)
}
