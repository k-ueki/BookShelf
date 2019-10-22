package repository

import (
	"database/sql"

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

func Register(db *sqlx.DB, user model.UserReq) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO users (firebase_uid,name,photo_url) VALUES (?,?,?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(user.Uid, user.Name, "")
}

func IsExistsUserByUid(db *sqlx.DB, uid string) (*model.User, error) {
	usr := model.User{}
	if err := db.Get(&usr, `SELECT id,name FROM users WHERE firebase_uid=?`, uid); err != nil {
		return nil, err
	}
	return &usr, nil
}

func IsExistsByDispName(db *sqlx.DB, dname string) (*model.UserDispName, error) {
	resp := model.UserDispName{}
	if err := db.Get(&resp, `SELECT user_id,disp_name FROM user_id WHERE disp_name=?`, dname); err != nil {
		return nil, err
	}

	return &resp, nil
}

func RegisterUid(db *sqlx.DB, id int, name string) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO user_id (user_id,disp_name) VALUES (?,?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id, name)
}

func main() {}
