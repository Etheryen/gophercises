package utils

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type PathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url"  json:"url"`
}

func ParseYaml(yml []byte) ([]PathUrl, error) {
	var parsedYaml []PathUrl

	err := yaml.Unmarshal(yml, &parsedYaml)
	if err != nil {
		return nil, err
	}

	return parsedYaml, nil
}

func BuildMap(pathsUrls []PathUrl) map[string]string {
	pathMap := make(map[string]string)

	for _, pathUrl := range pathsUrls {
		pathMap[pathUrl.Path] = pathUrl.URL
	}

	return pathMap
}

func ParseJSON(jsonInput []byte) ([]PathUrl, error) {
	var parsedJSON []PathUrl

	err := json.Unmarshal(jsonInput, &parsedJSON)
	if err != nil {
		return nil, err
	}

	return parsedJSON, nil
}
