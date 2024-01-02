package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"github.com/prioarief/gofiber-repository-pattern/models/converter"
	"github.com/prioarief/gofiber-repository-pattern/repositories"
)

// type BookService interface {
// 	GetBooks() ([]entities.Book, error)
// 	GetBookById(id int) (entities.Book, error)
// }

// type bookService struct {
// 	repository *repositories.BookRepository
// }

type BookService struct {
	Repository *repositories.BookRepository
	Validate   *validator.Validate
}

func NewBookService(r *repositories.BookRepository, validate *validator.Validate) *BookService {
	return &BookService{Repository: r, Validate: validate}
}

func (s *BookService) GetBooks() ([]models.BookResponse, error) {
	books, err := s.Repository.GetBooks()
	if err != nil {
		return nil, err
	}

	newBooks := make([]models.BookResponse, len(books))
	for i, book := range books {
		newBooks[i] = *converter.BookConverter(&book)
	}

	return newBooks, nil
}

func (s *BookService) GetBookById(id int) (*models.BookResponse, error) {
	book, err := s.Repository.GetBookById(id)
	if err != nil {
		return &models.BookResponse{}, err
	}

	return converter.BookConverter(&book), nil
}

func (s *BookService) Create(request *models.BookRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	err := s.Repository.Create(request)
	if err != nil {
		return err
	}

	return nil
}
