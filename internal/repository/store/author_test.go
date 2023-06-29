package store

import (
	"github.com/Victor1995fed/golang-test-cvado/config"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	bookIdSuccess = 1
	bookIdFail    = 999
	envPath       = "../../../.env"
)

// Тест метода получения автора по книге (успешный)
func TestGetAuthorsByBookSuccess(t *testing.T) {
	cfg := config.ParseConfig(envPath)
	db := database.Connect(&cfg)
	authors, _ := GetAuthorsByBook(database.New(db), bookIdSuccess)
	assert.NotEmpty(t, authors)
}

// Тест метода получения автора по книге (неудачный)
func TestGetAuthorsByBookNotFound(t *testing.T) {
	cfg := config.ParseConfig(envPath)
	db := database.Connect(&cfg)
	authors, _ := GetAuthorsByBook(database.New(db), bookIdFail)
	assert.Empty(t, authors)
}
