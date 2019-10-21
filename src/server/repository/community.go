package repository

import (
	"database/sql"

	"github.com/gocraft/dbr"
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

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
