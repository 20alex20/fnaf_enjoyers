package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Repository struct {
		DB struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Name     string `yaml:"name"`
		} `yaml:"db"`
	} `yaml:"repository"`
}

func ReadConfig(configPath string) (cfg *Config, err error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err = d.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
