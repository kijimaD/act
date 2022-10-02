package config

import (
	"github.com/go-yaml/yaml"
	"os"
)

var defaultConfigFilePath = ".act.yml"

// memo: 各フィールドはpublicである必要がある
type Config struct {
	User     string `yaml:"user"`
	OutType  string `yaml:"outType"`
	OutPath  string `yaml:"outPath"`
	IsCommit bool   `yaml:"commit"`
	IsPush   bool   `yaml:"push"`
}

func NewConfig() Config {
	// default parameter
	return Config{
		User: "",
		OutType: "stdout",
		OutPath: "./README.md",
		IsCommit:  false,
		IsPush:    false,
	}
}

func (c *Config) Load() {
	f, err := os.Open(defaultConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&c)
}
