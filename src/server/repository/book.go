package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

func GetBook(db *sqlx.DB, isbn string) (*model.Book, error) {
	resp := model.Book{}
	if err := db.Get(&resp, `SELECT * FROM books WHERE isbn=?`, isbn); err != nil {
		return nil, err
	}
	return &resp, nil
}

func SelectBookByUserId(db *sqlx.DB, user_id int64) (*[]model.Book, error) {
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

func Insert(db *sqlx.DB, book model.Book) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO books (title,author,price,img_url,page_url,isbn) VALUES (?,?,?,?,?,?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(book.Title, book.Author, book.Price, book.ImgUrl, book.PageUrl, book.Isbn)
}

func RegisterBookAndUser(db *sqlx.DB, user_id, book_id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO user_book (user_id,book_id) VALUES (?,?)
	`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(user_id, book_id)

}
