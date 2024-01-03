package configs

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      false,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2 Jan 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Format:     "[${time}] ${status} - ${method} ${path} body: ${body} queryParams: ${queryParams}\n",
	}))

	// app.Get("/monitor", monitor.New())

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return c.Status(code).JSON(fiber.Map{
			"code":    0,
			"message": err.Error(),
			"data":    nil,
		})
	}
}
