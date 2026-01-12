package utils

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	EmailHostUser     string
	EmailHostPassword string
}

func ReadConfiguration() (Configuration, error) {
	// get config from env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	// get config from os variable
	viper.AutomaticEnv()

	return Configuration{
		EmailHostUser:     viper.GetString("EMAIL_HOST_USER"),
		EmailHostPassword: viper.GetString("EMAIL_HOST_PASSWORD"),
	}, nil

}
