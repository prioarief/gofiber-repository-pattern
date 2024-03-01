package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"github.com/prioarief/gofiber-repository-pattern/services"
	"github.com/sirupsen/logrus"
)

type BookHandler struct {
	Service *services.BookService
	Log     *logrus.Logger
}

func NewBookHandler(s *services.BookService, log *logrus.Logger) *BookHandler {
	return &BookHandler{
		Service: s,
		Log:     log,
	}
}

func (b *BookHandler) List(c *fiber.Ctx) error {
	categoryId := c.QueryInt("category_id")
	keyword := c.Query("keyword")

	filter := &models.BookFilter{
		CategoryId: &categoryId,
		Keyword:    &keyword,
	}

	if err := c.QueryParser(filter); err != nil {
		b.Log.WithError(err).Error("failed to process request")
		return fiber.ErrBadRequest
	}

	books, err := b.Service.List(c.UserContext(), filter)
	if err != nil {
		return err
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
		b.Log.WithError(err).Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	book, err := b.Service.Get(c.UserContext(), id)
	if err != nil {
		return fiber.ErrNotFound
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
		b.Log.WithError(err).Error("failed to process request")
		return fiber.ErrBadRequest
		// return &fiber.Error{Message: "Opppss", Code: 400}
	}

	err := b.Service.Create(c.UserContext(), request)
	if err != nil {
		// return fiber.ErrBadRequest
		return &fiber.Error{Message: err.Error(), Code: 400}
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
		b.Log.WithError(err).Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	_, err = b.Service.Get(c.UserContext(), id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if err := b.Service.Delete(c.UserContext(), id); err != nil {
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
		b.Log.WithError(err).Error("failed parse param id")
		return fiber.ErrBadRequest
	}

	request := new(models.BookRequest)
	if err := c.BodyParser(request); err != nil {
		b.Log.WithError(err).Error("failed to process request")
		return fiber.ErrBadRequest
	}

	if err := b.Service.Update(c.UserContext(), id, request); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    1,
		"message": "success Update a book",
		"data":    request,
	})
}
