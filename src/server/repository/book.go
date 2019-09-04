package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

func SelectBookByUserId(db *sqlx.DB, user_id int) (*[]model.Book, error) {
	books := make([]model.Book, 0)
	if err := db.Select(&books, `
SELECT 
	books.id,
	books.title,
	books.author,
	books.price,
	books.img_url,
	books.page_url
FROM books
INNER JOIN user_book
ON user_book.book_id=books.id
WHERE user_book.user_id=?
	`, user_id); err != nil {
		return nil, err
	}
	return &books, nil
}
