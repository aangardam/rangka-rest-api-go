package config

import "github.com/spf13/viper"

type Config struct {
	BD_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_URL      string
	PORT        string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
