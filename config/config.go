package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username  string `yaml:"username"`
	Token     string `yaml:"token"`
	BaseURL   string `yaml:"baseURL"`
	BrowseURL string `yaml:"browseURL"`
}

func LoadConfig() (*Config, error) {
	conf, err := ioutil.ReadFile("./config/setting.yaml")
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
