package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	if err := godotenv.Load(); err != nil {
		config.AutomaticEnv()
	} else {
		config.SetConfigName(".env")
		config.SetConfigType("env")
		config.AddConfigPath(".")

		err := config.ReadInConfig()
		if err != nil {
			log.Fatalf("Error read config: %v", err)
		}
	}

	return config
}
