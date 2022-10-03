package config

import (
	"act/gh"
	"github.com/go-yaml/yaml"
	"os"
)

var defaultConfigFilePath = ".act.yml"

// memo: 各フィールドはpublicである必要がある
type Config struct {
	User User
	OutType  string `yaml:"outType"`
	OutPath  string `yaml:"outPath"`
	IsCommit bool   `yaml:"commit"`
	IsPush   bool   `yaml:"push"`
}

// Authenticated User
type User struct {
	Id    string
	Name  string
	Email string
}

func NewConfig() Config {
	// default parameter
	return Config{
		User: User{
			Id: "",
			Name: "",
			Email: "",
		},
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

	user := gh.New().User()
	c.User = User{
		Id: *user.Login,
		Name: *user.Name,
		Email: *user.Email,
	}
}
