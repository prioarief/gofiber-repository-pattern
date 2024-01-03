package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"github.com/prioarief/gofiber-repository-pattern/models/converter"
	"github.com/prioarief/gofiber-repository-pattern/repositories"
	"github.com/sirupsen/logrus"
)

// type BookService interface {
// 	List() ([]entities.Book, error)
// 	Get(id int) (entities.Book, error)
// }

// type bookService struct {
// 	repository *repositories.BookRepository
// }

type BookService struct {
	Repository *repositories.BookRepository
	Validate   *validator.Validate
	Log        *logrus.Logger
}

func NewBookService(r *repositories.BookRepository, validate *validator.Validate, log *logrus.Logger) *BookService {
	return &BookService{Repository: r, Validate: validate, Log: log}
}

func (s *BookService) List(ctx context.Context) ([]models.BookResponse, error) {
	books, err := s.Repository.List(ctx)
	if err != nil {
		return nil, err
	}

	newBooks := make([]models.BookResponse, len(books))
	for i, book := range books {
		newBooks[i] = *converter.BookConverter(&book)
	}

	return newBooks, nil
}

func (s *BookService) Get(ctx context.Context, id int) (*models.BookResponse, error) {
	book, err := s.Repository.Get(ctx, id)
	if err != nil {
		return &models.BookResponse{}, err
	}

	return converter.BookConverter(&book), nil
}

func (s *BookService) Create(ctx context.Context, request *models.BookRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	err := s.Repository.Create(ctx, request)
	if err != nil {
		s.Log.WithError(err).Error("failed to insert new data")
		return err
	}

	return nil
}

func (s *BookService) Delete(ctx context.Context, id int) error {
	err := s.Repository.Delete(ctx, id)
	if err != nil {
		s.Log.WithError(err).Error("failed to delete data")
		return err
	}

	return nil
}

func (s *BookService) Update(ctx context.Context, id int, request *models.BookRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	err := s.Repository.Update(ctx, id, request)
	if err != nil {
		s.Log.WithError(err).Error("failed to update data")
		return err
	}

	return nil
}
