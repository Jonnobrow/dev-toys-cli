package commands

type Command interface {
	Name() string
	Exec(string) (string, error)
}

type base struct {
	name string
	desc string
}

func NewBase(name, desc string) base {
	return base{
		name: name,
		desc: desc,
	}
}
