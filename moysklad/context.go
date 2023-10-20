package moysklad

type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}

func (m Context) String() string {
	return Stringify(m)
}
