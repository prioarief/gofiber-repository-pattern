package entities

type Book struct {
	Id          int      `gorm:"column:id;primaryKey"`
	Title       string   `gorm:"column:title"`
	Description string   `gorm:"column:description"`
	Price       int      `gorm:"column:price"`
	CategoryId  int      `gorm:"column:category_id"`
	Category    Category `gorm:"foreignKey:category_id;references:id"`
}

func (b *Book) TableName() string {
	return "books"
}
