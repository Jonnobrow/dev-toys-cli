package commands

type Command interface {
	Name() string
	Exec(string) (string, error)
	DisplayInput(string) string
	DisplayOutput(string) string
}

type base struct {
	name string
	desc string
}

func (b base) Name() string {
	return b.name
}

func (b base) Desc() string {
	return b.desc
}

func (b base) DisplayInput(input string) string {
	return input
}

func (b base) DisplayOutput(input string) string {
	return input
}

func NewBase(name, desc string) base {
	return base{
		name: name,
		desc: desc,
	}
}
