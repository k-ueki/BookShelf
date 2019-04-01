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
	for i, v := range tmp {
		tmptmp := strings.Split(v, "=")
		el[i] = tmptmp[1]
	}
	fmt.Println(el)
	res := map[string]string{
		"name":     el[0],
		"password": el[1],
		"email":    el[2],
	}
	return res
}
func (u *User) NewUser(w http.ResponseWriter, r *http.Request) {
	body := body(r)
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
