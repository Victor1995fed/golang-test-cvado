package store

import (
	"github.com/Victor1995fed/golang-test-cvado/internal/model"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
)

// GetAuthorsByBook Метод для получения авторов по id книги
func GetAuthorsByBook(repo *database.Repo, id int) ([]model.Author, error) {
	var authors []model.Author
	err := repo.Db.Model(&model.Author{}).
		Joins("JOIN author_book ON author.id = author_book.author_id").
		Where("author_book.book_id = ?", id).
		Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}
