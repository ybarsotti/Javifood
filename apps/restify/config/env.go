package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var EnvConfig Env

type Env struct {
	AppEnv string   `mapstructure:"APP_ENV"`
	DB     DbConfig `mapstructure:"DB"`
}

type DbConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Port     string `mapstructure:"PORT"`
	DB       string `mapstructure:"DB"`
	SslMode  string `mapstructure:"SSL_MODE"`
	TimeZone string `mapstructure:"TIME_ZONE"`
}

func NewEnv() *Env {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read env: ", err)
	}
	var env Env
	if err = v.Unmarshal(&env); err != nil {
		log.Fatal("env cannot be loaded: ", err)
	}

	if env.AppEnv == "dev" {
		log.Println("App is in dev mode")
	}

	EnvConfig = env
	return &env
}
