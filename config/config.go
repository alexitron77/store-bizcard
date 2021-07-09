package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Postgres db `mapstructure:"POSTGRES"`
	Mongo    db `mapstructure:"MONGO"`
}

type db struct {
	Url      string `mapstructure:"URL"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

func LoadConfig(path string) *config {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error opening config file")
	}

	conf := &config{}
	err = viper.Unmarshal(&conf)

	if err != nil {
		log.Fatal("Error unmarshal config file")
	}
	return conf
}
