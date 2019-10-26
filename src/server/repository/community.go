package repository

import (
	"database/sql"

	"github.com/gocraft/dbr"
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

func GetInfoByCid(db *sqlx.DB, cid int) (*model.Community, error) {
	com := model.Community{}
	if err := db.Get(&com, `SELECT * FROM communities WHERE id=?`, cid); err != nil {
		return nil, err
	}
	return &com, nil
}

func GetBooksByCid(db *sqlx.DB, cid int) (*[]model.Book, error) {
	books := make([]model.Book, 0)
	if err := db.Select(&books, `
SELECT DISTINCT
	books.id,
	books.title,
	books.author,
	books.price,
	books.img_url,
	books.page_url,
	books.isbn
FROM books 
INNER JOIN user_book ON user_book.book_id = books.id 
INNER JOIN user_community ON user_community.user_id = user_book.user_id 
WHERE user_community.community_id = ?;
	`, cid); err != nil {
		return nil, err
	}
	return &books, nil
}

func GetUsersByCid(db *sqlx.DB, cid int) (*[]model.User, error) {
	users := make([]model.User, 0)
	if err := db.Select(&users, `
SELECT DISTINCT
	users.id,
	users.firebase_uid,
	users.name,
	users.photo_url
FROM users 
INNER JOIN user_community ON user_community.user_id = users.id 
WHERE user_community.community_id = ?;
	`, cid); err != nil {
		return nil, err
	}
	return &users, nil
}

func GetAllByUid(db *sqlx.DB, uid string) (*[]model.Community, error) {
	communities := make([]model.Community, 0)
	if err := db.Select(&communities, `
SELECT DISTINCT 
	communities.id,
	communities.name
FROM communities
INNER JOIN user_community
ON user_community.community_id=communities.id
INNER JOIN users
ON user_community.user_id=users.id
WHERE users.firebase_uid=?
	`, uid); err != nil {
		return nil, err
	}
	return &communities, nil
}

func CreateCommunity(db *sqlx.DB, sess *dbr.Session, coms []model.CommunityParams) (sql.Result, error) {
	stmt := sess.InsertInto("user_community").Columns("user_id", "community_id")
	for _, v := range coms {
		stmt.Record(v)
	}
	return stmt.Exec()
}

func CreateNewCommunity(db *sqlx.DB, req *model.Community) (sql.Result, error) {

	stmt, err := db.Prepare(`
INSERT INTO communities (name) VALUES (?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(req.Name)
}

// func CreateCommunity(db *sqlx.DB, sess *dbr.Session, coms []model.CommunityParams) (sql.Result, error) {
// 	stmt := sess.InsertInto("user_community").Columns("user_id", "community_id")
// 	for _, v := range coms {
// 		stmt.Record(v)
// 	}
// 	return stmt.Exec()
// }
