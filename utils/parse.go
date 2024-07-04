package utils

import (
	"encoding/json"
	"goscrape/structs"
	"net/url"
	"os"
)

func ParseSiteHost(site string) (string, error) {
	url, err := url.Parse(site)
	if err != nil {
		return "", err
	}

	return url.Host, nil
}

func ParseConfig(filename string, configjson string) (*structs.Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if configjson == "" {
			return nil, err
		}
		file = []byte(configjson)
	}

	config := structs.Config{}
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
