package database

import (
	"gorm.io/gorm"
)

// Repo Структура для хранения данных
type Repo struct {
	Db *gorm.DB
}

// New Создание экземляра Repo
func New(db *gorm.DB) *Repo {
	return &Repo{Db: db}
}
