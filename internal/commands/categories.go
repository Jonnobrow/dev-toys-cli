package commands

type Category struct {
	Title       string
	Prompt      string
	Subcommands []Command
	Cursor      int
}

func NewCategory(title, prompt string, subcommands ...Command) *Category {
	return &Category{
		Title:       title,
		Prompt:      prompt,
		Subcommands: subcommands,
	}
}

func (c *Category) Selected() Command {
	return c.Subcommands[c.Cursor]
}

func (c *Category) CursorUp() {
	if c.Cursor > 0 {
		c.Cursor--
	}
}

func (c *Category) CursorDown() {
	if c.Cursor < len(c.Subcommands)-1 {
		c.Cursor++
	}
}
