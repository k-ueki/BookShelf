package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/k-ueki/app2/src/server/model"
	_ "github.com/k-ueki/app2/src/server/model"
)

type User struct {
	DB     *sql.DB
	Stream chan *model.User
}

func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	var usr model.User

	inserted, err := usr.Insert(u.DB)
	if err != nil {
		fmt.Println(inserted)
		fmt.Println("FAILED!!!", err)
		return
	}
	fmt.Println("SUCCESS", inserted)
	fmt.Println(r)
}
