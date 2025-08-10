package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort          int    `required:"true" split_words:"true"`
	BasePath            string `required:"true" split_words:"true"`
	DBHost              string `required:"true" split_words:"true"`
	DBPort              int    `required:"true" split_words:"true"`
	DBUser              string `required:"true" split_words:"true"`
	DBPassword          string `required:"true" split_words:"true"`
	DBName              string `required:"true" split_words:"true"`
	SSLMode             string `required:"true" split_words:"true"`
	SearchPath          string `required:"true" split_words:"true"`
	PoolMaxConns        int    `required:"true" split_words:"true"`
	PoolMinConns        int    `required:"true" split_words:"true"`
	PoolMaxConnIdleTime string `required:"true" split_words:"true"`
	ClientID            string `required:"true" split_words:"true"`
	ClientSecret        string `required:"true" split_words:"true"`
	RedirectURL         string `required:"true" split_words:"true"`
}

func LoadConfig() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println(err)
	}

	return &cfg
}
