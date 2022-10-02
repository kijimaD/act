package main

import (
	"github.com/go-yaml/yaml"
	"os"
)

var DefaultConfigFilePath = ".act.yml"

// memo: 各フィールドはpublicである必要がある
type Config struct {
	OutType string `yaml:"outType"`
}

func newConfig() Config {
	return Config{
		OutType: "stdout",
	}
}

func (c *Config) load(loadPath string) {
	f, err := os.Open(loadPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&c)
}
