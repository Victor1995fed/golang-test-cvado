package model

// Author Модель автора
type Author struct {
	Id    int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Books []Book `gorm:"many2many:author_book;"`
}

// TableName Возвращает название таблицы
func (Author) TableName() string {
	return "author"
}
