package models

type BookResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	CategoryName string `json:"category_name"`
}

type BookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required,number,gte=1000"`
}

type BookFilter struct {
	CategoryId *int    `json:"category_id" validate:"omitempty,gte=0"`
	Keyword    *string `json:"keyword" validate:"omitempty"`
}
