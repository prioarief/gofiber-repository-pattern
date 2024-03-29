package converter

import (
	"github.com/prioarief/gofiber-repository-pattern/entities"
	"github.com/prioarief/gofiber-repository-pattern/models"
)

func BookConverter(book *entities.Book) *models.BookResponse {
	return &models.BookResponse{
		ID:           book.Id,
		Title:        book.Title,
		Description:  book.Description,
		Price:        book.Price,
		CategoryName: book.Category.Name,
	}
}
