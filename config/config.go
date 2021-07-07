package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	DBUrl      string `mapstructure:"DB_URL"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
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
