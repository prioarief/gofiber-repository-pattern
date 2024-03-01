package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/prioarief/gofiber-repository-pattern/entities"
	"github.com/prioarief/gofiber-repository-pattern/helpers"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"github.com/prioarief/gofiber-repository-pattern/models/converter"
	"github.com/prioarief/gofiber-repository-pattern/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BookService struct {
	Repository *repositories.BookRepository
	Validate   *validator.Validate
	Log        *logrus.Logger
	DB         *gorm.DB
}

func NewBookService(r *repositories.BookRepository, validate *validator.Validate, log *logrus.Logger, db *gorm.DB) *BookService {
	return &BookService{Repository: r, Validate: validate, Log: log, DB: db}
}

func (s *BookService) List(ctx context.Context, filter *models.BookFilter) ([]models.BookResponse, error) {
	tx := s.DB.WithContext(ctx)
	defer tx.Rollback()

	if err := helpers.ValidationError(s.Validate, filter); err != nil {
		s.Log.WithError(err).Error("failed to validate request query params")
		return nil, err
	}

	books, err := s.Repository.List(tx, filter)
	if err != nil {
		s.Log.WithError(err).Error("failed get book lists")
		return nil, err
	}

	newBooks := make([]models.BookResponse, len(books))
	for i, book := range books {
		newBooks[i] = *converter.BookConverter(&book)
	}

	return newBooks, nil
}

func (s *BookService) Get(ctx context.Context, id int) (*models.BookResponse, error) {
	tx := s.DB.WithContext(ctx)
	defer tx.Rollback()

	book := new(entities.Book)

	err := s.Repository.Get(tx, id, book)
	if err != nil {
		s.Log.WithError(err).Error("failed get book detail")
		return nil, err
	}

	return converter.BookConverter(book), nil
}

func (s *BookService) Create(ctx context.Context, request *models.BookRequest) error {
	// if err := s.Validate.Struct(request); err != nil {
	if err := helpers.ValidationError(s.Validate, request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	tx := s.DB.WithContext(ctx)
	defer tx.Rollback()

	newRequest := &entities.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
	}

	err := s.Repository.Create(tx, newRequest)
	if err != nil {
		s.Log.WithError(err).Error("failed to insert new data")
		return err
	}

	return nil
}

func (s *BookService) Delete(ctx context.Context, id int) error {
	tx := s.DB.WithContext(ctx)
	defer tx.Rollback()

	book := new(entities.Book)
	err := s.Repository.Get(tx, id, book)
	if err != nil {
		s.Log.WithError(err).Error("failed get book detail")
		return err
	}

	err = s.Repository.Delete(tx, book)
	if err != nil {
		s.Log.WithError(err).Error("failed to delete data")
		return err
	}

	return nil
}

func (s *BookService) Update(ctx context.Context, id int, request *models.BookRequest) error {
	// if err := s.Validate.Struct(request); err != nil {
	if err := helpers.ValidationError(s.Validate, request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return err
	}

	tx := s.DB.WithContext(ctx)
	defer tx.Rollback()

	book := new(entities.Book)
	err := s.Repository.Get(tx, id, book)
	if err != nil {
		s.Log.WithError(err).Error("failed get book detail")
		return err
	}

	book.Title = request.Title
	book.Description = request.Description
	book.Price = request.Price

	err = s.Repository.Update(tx, id, book)
	if err != nil {
		s.Log.WithError(err).Error("failed to update data")
		return err
	}

	return nil
}
