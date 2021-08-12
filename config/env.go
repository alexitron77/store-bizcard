package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Postgres db    `mapstructure:"POSTGRES"`
	Mongo    db    `mapstructure:"MONGO"`
	Aws      aws   `mapstructure:"AWS"`
	Kafka    kafka `mapstructure:"KAFKA"`
}

type aws struct {
	AccessKey string `mapstructure:"ACCESS-KEY"`
	Secret    string `mapstructure:"SECRET"`
}

type db struct {
	Url      string `mapstructure:"URL"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type kafka struct {
	Url   string `mapstructure:"URL"`
	Topic string `mapstructure:"TOPIC"`
}

func LoadConfig(path string) *Env {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error opening config file")
	}

	conf := &Env{}
	err = viper.Unmarshal(&conf)

	if err != nil {
		log.Fatal("Error unmarshal config file")
	}
	return conf
}
