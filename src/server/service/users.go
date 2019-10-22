package service

import (
	"fmt"

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
func (u *User) Index(uid string) (interface{}, error) {
	res := model.UserResp{}

	usr, err := repository.SelectUserByUid(u.DB, uid)
	if err != nil {
		return nil, err
	}
	fmt.Println("user", usr)

	books, err := repository.SelectBookByUserId(u.DB, usr.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("books", books)

	coms, err := repository.GetAllByUid(u.DB, uid)
	if err != nil {
		return nil, err
	}
	fmt.Println("coms,", coms)

	res.Id = usr.Id
	res.Name = usr.Name
	res.DispName = usr.DispName
	res.Books = *books
	res.Communities = *coms

	return res, nil
}

func (u *User) IsExists(uid string) (*int64, bool) {
	user, err := repository.IsExistsUserByUid(u.DB, uid)
	if err != nil {
		return nil, false
	}
	return &user.Id, true
}

func (u *User) IsOKUserId(disp_name string) (bool, error) {
	info, _ := repository.IsExistsByDispName(u.DB, disp_name)
	if info != nil {
		return false, nil
	}

	return true, nil
}

func (u *User) RegisterUserId(Uid int, dispname string) error {
	_, err := repository.RegisterUid(u.DB, Uid, dispname)
	if err != nil {
		return err
	}
	return nil
}
func main() {}
