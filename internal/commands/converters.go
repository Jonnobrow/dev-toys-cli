package commands

import (
	"bytes"
	"encoding/json"

	"github.com/ghodss/yaml"
)

var (
	JSONtoYAML = jsonToYaml{
		base: NewBase("JSON -> YAML", "Convert JSON to YAML"),
	}
	YAMLtoJSON = yamlToJson{
		base: NewBase("YAML -> JSON", "Convert YAML to JSON"),
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
	return prettyJson(string(json)), nil
}

func prettyJson(rawJSON string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(rawJSON), "", "  "); err != nil {
		return ""
	}
	return prettyJSON.String()
}
