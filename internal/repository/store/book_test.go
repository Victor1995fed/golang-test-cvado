package store

import (
	"github.com/Victor1995fed/golang-test-cvado/config"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	authorIdSuccess = 1
	authorIdFail    = 999
)

// Тест метода получения автора по книге(успешный)
func TestGetBooksByAuthorSuccess(t *testing.T) {
	cfg := config.ParseConfig(envPath)
	db := database.Connect(&cfg)
	authors, _ := GetBooksByAuthor(database.New(db), authorIdSuccess)
	assert.NotEmpty(t, authors)
}

// Тест метода получения автора по книге (не найден)
func TestGetBooksByAuthorNotFound(t *testing.T) {
	cfg := config.ParseConfig(envPath)
	db := database.Connect(&cfg)
	authors, _ := GetBooksByAuthor(database.New(db), authorIdFail)
	assert.Empty(t, authors)
}
