package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/k-ueki/app2/src/server/model"
	_ "github.com/k-ueki/app2/src/server/model"
)

type User struct {
	DB     *sql.DB
	Stream chan *model.User
}

func body(r *http.Request) map[string]string {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println(string(body))
	tmp := string(body)
	return sep(tmp, "&")
}
func sep(str string, cha string) map[string]string {
	tmp := strings.Split(str, cha)
	fmt.Println(tmp)

	var el = make([]string, 3) //name,password,emailの3
	var th = make([]string, 3)
	for i, v := range tmp {
		tmptmp := strings.Split(v, "=")
		th[i] = tmptmp[0]
		el[i] = tmptmp[1]
	}
	fmt.Println(el)
	res := map[string]string{
		th[0]: el[0],
		th[1]: el[1],
		th[2]: el[2],
	}
	return res
}

//新規登録
func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	var usr = model.User{
		Name:     body["name"],
		Password: body["password"],
		Email:    body["email"],
	}
	inserted, err := usr.Insert(u.DB)
	if err != nil {
		fmt.Println(inserted)
		fmt.Println("FAILED!!!", err)
		return
	}
	fmt.Println("SUCCESS", inserted)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	var usr = model.User{
		Password: body["pass"],
		Email:    body["email"],
	}
	res, err := usr.Check(u.DB)
	if err != nil {
		fmt.Println("ERROR!", err)
		return
	}
	fmt.Println("CHECKED", res)
}
