package main

import (
	"github.com/go-yaml/yaml"
	"os"
)

var DefaultConfigFilePath = ".act.yml"

// memo: 各フィールドはpublicである必要がある
type Config struct {
	OutType  string `yaml:"outType"`
	OutPath  string `yaml:"outPath"`
	IsCommit bool   `yaml:"commit"`
	IsPush   bool   `yaml:"push"`
}

func newConfig() Config {
	// default parameter
	return Config{
		OutType: "stdout",
		OutPath: "./README.md",
		IsCommit:  false,
		IsPush:    false,
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
