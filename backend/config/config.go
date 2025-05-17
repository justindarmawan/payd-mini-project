package config

import (
	"log"

	"github.com/spf13/viper"
)

var JWTSecret string

func InitConfig() {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	JWTSecret = viper.GetString("jwt.secret")
	if JWTSecret == "" {
		log.Fatal("JWT secret is missing in config")
	}
}
