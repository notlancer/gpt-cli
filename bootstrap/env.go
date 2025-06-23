package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	BearerToken string `mapstructure:"BearerToken"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment can't be loaded: ", err)
	}

	return &env
}
