package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host          string `yaml:"HTTP_HOST"`      //nolint: tagliatelle
	Port          string `yaml:"HTTP_PORT"`      //nolint: tagliatelle
	CacheCapacity int    `yaml:"CACHE_CAPACITY"` //nolint: tagliatelle
}

func Configure() (*Config, error) {
	confContent, err := os.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err = yaml.Unmarshal(confContent, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
