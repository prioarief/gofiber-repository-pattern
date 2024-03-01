package repositories

import (
	"github.com/prioarief/gofiber-repository-pattern/entities"
	"github.com/prioarief/gofiber-repository-pattern/models"
	"gorm.io/gorm"
)

// type BookRepository interface {
// 	List() ([]entities.Book, error)
// 	Get(id int) (entities.Book, error)
// }

// type bookRepository struct {
// 	db *gorm.DB
// }

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) List(db *gorm.DB, filter *models.BookFilter) ([]entities.Book, error) {
	var books []entities.Book
	query := db.Preload("Category")

	if *filter.CategoryId != 0 {
		query.Where("category_id = ?", &filter.CategoryId)
	}

	if *filter.Keyword != "" {
		query.Where("title LIKE ?", "%"+*filter.Keyword+"%")
	}

	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Get(db *gorm.DB, id int, book *entities.Book) error {

	if err := db.Where("id = ?", id).Preload("Category").First(book).Error; err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Create(db *gorm.DB, request *entities.Book) error {
	return db.Create(request).Error
}

func (r *BookRepository) Update(db *gorm.DB, id int, request *entities.Book) error {

	return db.Save(request).Error
}

func (r *BookRepository) Delete(db *gorm.DB, book *entities.Book) error {
	return db.Delete(&book).Error
}
