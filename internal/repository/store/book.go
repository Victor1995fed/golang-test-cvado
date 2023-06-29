package store

import (
	"github.com/Victor1995fed/golang-test-cvado/internal/model"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
)

// GetBooksByAuthor Метод для получения книг по id автора
func GetBooksByAuthor(repo *database.Repo, id int) ([]model.Book, error) {
	var books []model.Book
	err := repo.Db.Model(&model.Book{}).
		Joins("JOIN author_book ON book.id = author_book.book_id").
		Where("author_book.author_id = ?", id).
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
