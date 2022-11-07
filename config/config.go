package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
	Db   map[string]string
}

func GetConfig() (*Config, error) {
	var config Config
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the Config file
	if err != nil {               // Handle errors reading the Config file
		panic(fmt.Errorf("Fatal error Config file: %s \n", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &config, nil
}
