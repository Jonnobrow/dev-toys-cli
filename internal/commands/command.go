package commands

type Command interface {
	Name() string
	Exec(string) (string, error)
	DisplayInput(string) string
	DisplayOutput(string) string
	ShouldDisplayInput() bool
}

type base struct {
	name string
	desc string

	displayInput  bool
	displayOutput bool
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

func (b base) ShouldDisplayInput() bool {
	return b.displayInput
}

func NewBase(name, desc string) base {
	newBase := base{
		name:          name,
		desc:          desc,
		displayInput:  true,
		displayOutput: true,
	}

	return newBase
}

func (b base) withoutInputDisplay() base {
	var newB base
	newB = b
	newB.displayInput = false
	return newB
}
