package configs

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/prioarief/gofiber-repository-pattern/handlers"
	"github.com/prioarief/gofiber-repository-pattern/repositories"
	"github.com/prioarief/gofiber-repository-pattern/routers"
	"github.com/prioarief/gofiber-repository-pattern/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB       *sql.DB
	App      *fiber.App
	Config   *viper.Viper
	Validate *validator.Validate
	Log      *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	bookRepository := repositories.NewBookRepository(config.DB)

	// setup services
	bookService := services.NewBookService(bookRepository, config.Validate, config.Log)

	// setup handler
	bookHandler := handlers.NewBookHandler(bookService, config.Log)

	// setup route
	routeConfig := routers.RouteConfig{
		App:         config.App,
		BookHandler: bookHandler,
	}

	routeConfig.Setup()
}
