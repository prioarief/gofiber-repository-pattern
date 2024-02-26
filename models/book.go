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
