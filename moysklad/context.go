package moysklad

type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}

func (context Context) String() string {
	return Stringify(context)
}
