package commands

import "strings"

type Category struct {
	Title       string
	Prompt      string
	Subcommands []Command
	Cursor      int
}

var (
	Categories = []*Category{
		// Converters
		NewCategory(
			"Converters", "Convert Stuff",
			Converters()...,
		),
		// Formatters
		NewCategory(
			"Formatters", "Format Stuff",
			Formatters()...,
		),
		// Generators
		NewCategory(
			"Generators", "Generate Stuff",
			Generators()...,
		),
	}
)

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

func (c *Category) CliName() string {
	return strings.ToLower(c.Title)
}

func (c *Category) CliShort() byte {
	return c.CliName()[0]
}
