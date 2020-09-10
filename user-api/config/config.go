package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
}

var Env *Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	viper.Unmarshal(&Env)
}
