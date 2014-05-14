package config

import (
	"encoding/json"
)

type Config struct {
	Path      string
	Pattern   string
	Ignore    string
	Command   string
	Recursive bool
}

func Parse(input []byte) ([]Config, error) {
	var configs []Config

	if err := json.Unmarshal(input, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}
