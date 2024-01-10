package repositories

import (
	"context"
	"database/sql"

	"github.com/prioarief/gofiber-repository-pattern/entities"
	"github.com/prioarief/gofiber-repository-pattern/models"
)

// type BookRepository interface {
// 	List() ([]entities.Book, error)
// 	Get(id int) (entities.Book, error)
// }

// type bookRepository struct {
// 	db *sql.DB
// }

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) List(ctx context.Context) ([]entities.Book, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) Get(ctx context.Context, id int) (*entities.Book, error) {
	var book entities.Book

	row := r.DB.QueryRow("SELECT * FROM books WHERE id = ?", id)
	if err := row.Scan(&book.Id, &book.Title, &book.Description, &book.Price); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) Create(ctx context.Context, request *models.BookRequest) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO books (title, description, price) VALUES (?,?, ?)", request.Title, request.Description, request.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Update(ctx context.Context, id int, request *models.BookRequest) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE books SET title = ?, description = ?, price = ? WHERE id = ?", request.Title, request.Description, request.Price, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
