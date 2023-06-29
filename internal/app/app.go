package app

import (
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
)

// App структура сервера grpc.
type App struct {
	cvadoProto.UnimplementedCvadoServer
	Repo *database.Repo
}

// NewServer Возвращает ссылку на структуру сервера
func NewServer(repo *database.Repo) *App {
	return &App{
		Repo: repo,
	}
}
