package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/repository"
)

type Book struct {
	DB *sqlx.DB
}

func NewBookService(db *sqlx.DB) *Book {
	return &Book{db}
}

func (b *Book) IsExists(isbn string) bool {
	book, _ := repository.GetBook(b.DB, isbn)
	if book == nil {
		return false
	}
	return true
}
