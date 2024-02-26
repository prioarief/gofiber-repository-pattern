package entities

type Category struct {
	Id   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	// Books []Book `gorm:"foreignKey:category_id;references:id"`
}

func (c *Category) TableName() string {
	return "categories"
}
