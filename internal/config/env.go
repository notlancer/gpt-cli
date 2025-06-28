package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	BearerToken string `mapstructure:"BearerToken"`
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GPT_CLI")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Warning: Could not read config file: %v", err)
	}

	env := &Env{}
	if err := viper.Unmarshal(env); err != nil {
		log.Fatalf("Environment can't be loaded: %v", err)
	}

	if env.BearerToken == "" {
		log.Fatal("BearerToken is required")
	}

	return env
}

func (e *Env) GetBearerToken() string {
	return e.BearerToken
}
