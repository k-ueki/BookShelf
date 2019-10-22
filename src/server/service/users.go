package service

import (
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

	books, err := repository.SelectBookByUserId(u.DB, usr.Id)
	if err != nil {
		return nil, err
	}

	coms, err := repository.GetAllByUid(u.DB, uid)
	if err != nil {
		return nil, err
	}

	res.Id = usr.Id
	res.Name = usr.Name
	res.Books = *books
	res.Communities = *coms

	dname, err := repository.SelectDispNameByUid(u.DB, uid)
	if err != nil {
		return res, nil
	}
	res.DispName = *dname

	return res, nil
}

func (u *User) IsExists(uid string) (*int64, bool) {
	user, err := repository.IsExistsUserByUid(u.DB, uid)
	if err != nil {
		return nil, false
	}
	return &user.Id, true
}

func (u *User) IsOKUserId(uid int, disp_name string) (bool, error) {
	info, _ := repository.GetDispInfoByUsrId(u.DB, uid)
	if info != nil {
		return false, nil
	}

	info, _ = repository.IsExistsByDispName(u.DB, disp_name)
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
