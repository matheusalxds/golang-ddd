package env

import (
	"log"

	"github.com/spf13/viper"
)

type EnvLoader interface {
	LoadEnv() *Env
}

type Env struct {
	Port    string
	DBHost  string
	DBPort  int
	DBUser  string
	DBPasss string
	DBName  string
}

func (*Env) LoadEnv() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading environments: %v", err)
	}

	return &Env{
		Port:    viper.GetString("PORT"),
		DBHost:  viper.GetString("DB_HOST"),
		DBPort:  viper.GetInt("DB_PORT"),
		DBUser:  viper.GetString("DB_USER"),
		DBPasss: viper.GetString("DB_PASS"),
		DBName:  viper.GetString("DB_NAME"),
	}
}

func NewEnv() EnvLoader {
	return &Env{}
}
