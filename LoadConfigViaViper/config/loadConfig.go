package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Greeting   string `json:"greeting"`
	Name       string `json:"name"`
	Number     int    `json:"number"`
	Executions int    `json:"executions"`
}

func Load(appName string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/" + appName + "/")   // look here first
	viper.AddConfigPath("$HOME/." + appName + "/") // then here
	viper.AddConfigPath("./config")                // finally here

	err = viper.ReadInConfig() // can file be read?
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config) // can file be unmarshaled?
	return
}
