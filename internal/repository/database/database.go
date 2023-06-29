package database

import (
	"fmt"
	"github.com/Victor1995fed/golang-test-cvado/config"
	"gorm.io/driver/mysql"
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

// Connect Подключение к базе
func Connect(cfg *config.Config) *gorm.DB {
	//Подключение БД
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
