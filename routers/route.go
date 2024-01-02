package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prioarief/gofiber-repository-pattern/handlers"
)

type RouteConfig struct {
	App         *fiber.App
	BookHandler *handlers.BookHandler
}

func (c *RouteConfig) Setup() {
	c.App.Get("/api/books", c.BookHandler.GetBooks)
	c.App.Get("/api/books/:id", c.BookHandler.GetBookById)
	c.App.Post("/api/books", c.BookHandler.Create)
}
