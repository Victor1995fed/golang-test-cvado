package app

import (
	"context"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/store"
	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
)

// GetBooksByAuthor Получает книги по id автора
func (s *App) GetBooksByAuthor(ctx context.Context, in *cvadoProto.RequestAuthor) (*cvadoProto.ResponseBook, error) {

	books, err := store.GetBooksByAuthor(s.Repo, int(in.GetId()))
	if err != nil {
		return nil, err
	}
	var list []*cvadoProto.Book
	for _, v := range books {
		list = append(list, &cvadoProto.Book{
			Id:    int64(v.Id),
			Title: v.Title,
		})
	}
	return &cvadoProto.ResponseBook{
		Books: list,
	}, nil
}

// GetAuthorByBook Получает автора по id книги
func (s *App) GetAuthorByBook(ctx context.Context, in *cvadoProto.RequestBook) (*cvadoProto.ResponseAuthor, error) {

	authors, err := store.GetAuthorsByBook(s.Repo, int(in.GetId()))
	if err != nil {
		return nil, err
	}
	var list []*cvadoProto.Author
	for _, v := range authors {
		list = append(list, &cvadoProto.Author{
			Id:   int64(v.Id),
			Name: v.Name,
		})
	}
	return &cvadoProto.ResponseAuthor{
		Authors: list,
	}, nil
}
