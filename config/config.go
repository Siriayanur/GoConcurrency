package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbName"`
	} `yaml:"db"`
}

func LoadConfig() Config {
	var cfg Config
	err := cleanenv.ReadConfig("application.yaml", &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfg
}
