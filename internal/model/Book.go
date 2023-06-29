package model

// Book Модель книги
type Book struct {
	Id      int      `gorm:"column:id"`
	Title   string   `gorm:"column:title"`
	Authors []Author `gorm:"many2many:author_book;"`
}

// TableName Возвращает название таблицы
func (Book) TableName() string {
	return "book"
}
