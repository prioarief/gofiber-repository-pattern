package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"github.com/prioarief/gofiber-repository-pattern/services"
)

type BookHandler struct {
	Service *services.BookService
}

func NewBookHandler(s *services.BookService) *BookHandler {
	return &BookHandler{
		Service: s,
	}
}

func (b *BookHandler) List(c *fiber.Ctx) error {
	books, err := b.Service.List()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "list of books",
		"data":    books,
	})
}

func (b *BookHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	book, err := b.Service.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		} else {
			return fiber.ErrInternalServerError
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "detail of book",
		"data":    book,
	})
}

func (b *BookHandler) Create(c *fiber.Ctx) error {
	request := new(models.BookRequest)

	if err := c.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
		// return &fiber.Error{Message: "Opppss", Code: 400}
	}

	err := b.Service.Create(request)
	if err != nil {
		return fiber.ErrBadRequest
		// return &fiber.Error{Message: err.Error(), Code: 400}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success insert new book",
		"data":    request,
	})
}

func (b *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	_, err = b.Service.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		} else {
			return fiber.ErrInternalServerError
		}
	}

	if err := b.Service.Delete(id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success delete a book",
		"data":    nil,
	})
}

func (b *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	_, err = b.Service.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrNotFound
		} else {
			return fiber.ErrInternalServerError
		}
	}

	request := new(models.BookRequest)
	if err := c.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := b.Service.Update(id, request); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success Update a book",
		"data":    request,
	})
}
