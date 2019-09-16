package service

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/repository"
)

type User struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *User {
	return &User{db}
}

// func (u *User) Index(uid string) (model.UserResp, error) {
func (u *User) Index(uid string) (int, interface{}, error) {
	res := model.UserResp{}

	usr, err := repository.SelectUserByUid(u.DB, uid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	books, err := repository.SelectBookByUserId(u.DB, usr.Id)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res.Id = usr.Id
	res.Name = usr.Name
	res.Books = *books

	// marUsr, _ := json.Marshal(res)

	// fmt.Fprintf(w, string(marUsr))
	return http.StatusOK, res, nil
}

func main() {}
