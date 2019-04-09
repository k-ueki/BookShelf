package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/k-ueki/app2/src/server/model"
	_ "github.com/k-ueki/app2/src/server/model"
)

type User struct {
	DB     *sql.DB
	Stream chan *model.User
}

//http.Requestからbodyをmapで返す
func body(r *http.Request) map[string]string {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	//fmt.Println(string(body))
	tmp := string(body)
	return sep(tmp, "&")
}
func sep(str string, cha string) map[string]string {
	tmp := strings.Split(str, cha)

	var el = make([]string, len(tmp)) //name,password,emailの3
	var th = make([]string, len(tmp))
	for i, v := range tmp {
		tmptmp := strings.Split(v, "=")
		th[i] = tmptmp[0]
		el[i] = tmptmp[1]
	}
	//fmt.Println(el)
	//res := map[string]string{
	//	th[0]: el[0],
	//	th[1]: el[1],
	//	th[2]: el[2],
	//}
	var res = make(map[string]string, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[th[i]] = el[i]
	}

	return res
}

//User新規登録
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
	//fmt.Println("SUCCESS", inserted)

	res, err := usr.SelectIdByNameANDPass(u.DB)
	if err != nil {
		fmt.Println("ERROR!", err)
		return
	}
	mar, _ := json.Marshal(res)
	fmt.Fprintf(w, string(mar))
}

//ログイン
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
	//fmt.Println("CHECKED", res)

	//response := fmt.Sprintf("%d,%s", res.ID, res.Name)
	fmt.Fprintf(w, fmt.Sprint(res.ID))
	//fmt.Fprintf(w, response)
}

//Personal Info をIDから取得
func (u *User) SelectPersonalInfo(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	//fmt.Println(body["id"])
	tmp, _ := strconv.Atoi(body["id"])
	var usr = model.User{
		ID: tmp,
	}
	//fmt.Println("USR", usr)
	usrInfo, err := usr.SelectById(u.DB)
	if err != nil {
		fmt.Println("Error!", err)
	}
	//fmt.Println("USRINFO", usrInfo)
	marUsr, _ := json.Marshal(usrInfo)
	fmt.Fprintf(w, string(marUsr))
}

//本の新規登録。楽天Books API を内部で叩く
func (u *User) GetBooksInfo(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	fmt.Println(body)

	applicationId := "1002996369552597120"
	title := body["booksTitle"]
	url := "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?"
	url += "applicationId=" + applicationId
	url += "&title=" + title

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bod, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(bod))
}
func (u *User) RegistBook(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	fmt.Println("BOBOBO", body)
}
