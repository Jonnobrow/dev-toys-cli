package commands

import (
	"bytes"
	"encoding/json"
)

var (
	Formatters = func() []Command {
		var commands []Command

		commands = append(commands, jsonPrettifier{
			base: NewBase("JSON Pretty Print", ""),
		})

		commands = append(commands, jsonMinifier{
			base: NewBase("JSON Minifier", ""),
		})

		return commands
	}
)

type jsonPrettifier struct{ base }

func (f jsonPrettifier) Exec(input string) (string, error) {
	return prettyJson(input)
}

type jsonMinifier struct{ base }

func (f jsonMinifier) Exec(input string) (string, error) {
	var miniJson bytes.Buffer
	if err := json.Compact(&miniJson, []byte(input)); err != nil {
		return "", err
	}
	return miniJson.String(), nil
}

func prettyJson(rawJSON string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(rawJSON), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
