package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
)

// isJSON determines if a given Config File Path is JSON
func isJSON(fileName string) bool {
	return strings.Contains(fileName, "json")
}

// IsYAML determines if a given Config File Path is YAML
func isYAML(fileName string) bool {
	return strings.Contains(fileName, "yml") || strings.Contains(fileName, "yaml")
}

// Load a new Config File from relative string path
func Load(configFilePath string) (*Config, error) {

	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, errors.New("File not found")
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Unable to read Config File")
	}

	var data *Config

	if isJSON(configFilePath) {
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return nil, errors.New("Unable to Parse config file content. This may be due to a format error")
		}
	}

	if isYAML(configFilePath) {
		err = yaml.Unmarshal(bytes, &data)
		if err != nil {
			return nil, errors.New("Unable to Parse config file content. This may be due to a format error")
		}
	}

	return data, nil
}
