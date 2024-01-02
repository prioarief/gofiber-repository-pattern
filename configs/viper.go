package configs

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName(".env")
	config.SetConfigType("env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error read config: %v", err)
	}

	return config
}
