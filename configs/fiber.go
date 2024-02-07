package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prioarief/gofiber-repository-pattern/helpers"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: helpers.NewErrorHandler(),
		Prefork:      false,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Format:     "[${time}] ${status} - ${method} ${path} body: ${body} queryParams: ${queryParams}\n",
	}))

	// app.Get("/monitor", monitor.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	return app
}
