package commands

import (
	"fmt"
	"strconv"

	"github.com/ghodss/yaml"
)

var (
	prettyBases = map[int]string{
		2:  "Binary",
		8:  "Octal",
		10: "Decimal",
		16: "Hexadecimal",
	}
	Converters = func() []Command {
		var converters []Command
		// YAML <> JSON
		converters = append(converters, yamlToJson{
			base: NewBase("YAML -> JSON", "Convert YAML to JSON").withAliases([]string{"y2j"}),
		})
		converters = append(converters, jsonToYaml{
			base: NewBase("JSON -> YAML", "Convert JSON to YAML").withAliases([]string{"j2y"}),
		})

		// Number Base Converters
		bases := []int{2, 8, 10, 16}
		for _, from := range bases {
			for _, to := range bases {
				if from != to {
					prettyTo := toPrettyBase(to)
					prettyFrom := toPrettyBase(from)
					converters = append(converters, numberBaseConverter{
						base:     NewBase(fmt.Sprintf("%s -> %s", prettyFrom, prettyTo), fmt.Sprintf("Convert %s to %s", prettyFrom, prettyTo)),
						fromBase: from,
						toBase:   to,
					})
				}
			}
		}
		return converters
	}
)

type jsonToYaml struct {
	base
}

func (c jsonToYaml) Exec(rawJSON string) (string, error) {
	yaml, err := yaml.JSONToYAML([]byte(rawJSON))
	if err != nil {
		return "", err
	}
	return string(yaml), nil
}

type yamlToJson struct {
	base
}

func (c yamlToJson) Exec(rawYAML string) (string, error) {
	json, err := yaml.YAMLToJSON([]byte(rawYAML))
	if err != nil {
		return "", err
	}
	return prettyJson(string(json))
}

type numberBaseConverter struct {
	base
	fromBase int
	toBase   int
}

func (c numberBaseConverter) Exec(rawDecimal string) (string, error) {
	return convertNumberBase(rawDecimal, c.fromBase, c.toBase)
}

func (c numberBaseConverter) DisplayInput(input string) string {
	return displayBasedNumber(input, c.fromBase)
}

func (c numberBaseConverter) DisplayOutput(input string) string {
	return displayBasedNumber(input, c.toBase)
}

func convertNumberBase(input string, fromBase int, toBase int) (string, error) {
	inputInt, err := strconv.ParseInt(input, fromBase, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(inputInt, toBase), nil
}

func displayBasedNumber(input string, base int) string {
	return fmt.Sprintf("%s (base %d)", input, base)
}

func toPrettyBase(base int) string {
	if pretty, ok := prettyBases[base]; ok {
		return pretty
	}
	return fmt.Sprintf("Base %d", base)
}
