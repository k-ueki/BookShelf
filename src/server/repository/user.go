package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
)

func SelectUserByUid(db *sqlx.DB, uid string) (*model.User, error) {
	usr := model.User{}
	if err := db.Get(&usr, `SELECT id,name,photo_url FROM users WHERE firebase_uid=?`, uid); err != nil {
		return nil, err
	}
	return &usr, nil
}

func main() {}
