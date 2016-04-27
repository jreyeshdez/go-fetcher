package config

import (
  "encoding/json"
	"io/ioutil"
)

type Config struct {
	Host  string
	OrgList []string
}

func Parse(configPath string) (*Config, error) {
  jsonBlob, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(jsonBlob, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
