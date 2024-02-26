package repositories

import (
	"github.com/prioarief/gofiber-repository-pattern/entities"
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

func (r *BookRepository) List(db *gorm.DB) ([]entities.Book, error) {
	var books []entities.Book
	// if err := db.Find(&books).Error; err != nil {
	if err := db.Preload("Category").Find(&books).Error; err != nil {
		return nil, err
	}

	// sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Preload("Category").Find(&books)
	// })

	// fmt.Println(sql)

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
