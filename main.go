package main

import (
	"fmt"

	"github.com/prioarief/gofiber-repository-pattern/configs"
)

func main() {
	viperConfig := configs.NewViper()
	app := configs.NewFiber(viperConfig)
	db := configs.NewDatabase(viperConfig)
	validate := configs.NewValidator(viperConfig)
	log := configs.NewLogger(viperConfig)

	configs.Bootstrap(&configs.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   viperConfig,
		Validate: validate,
		Log:      log,
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", viperConfig.GetInt("port"))))
}
