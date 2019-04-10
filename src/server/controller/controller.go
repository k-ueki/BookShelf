package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/k-ueki/app2/src/server/model"
	_ "github.com/k-ueki/app2/src/server/model"
)

type DBHandler struct {
	DB     *sql.DB
	Stream chan *model.User
}

//return用
type Res struct {
	ID    int
	Name  string
	Email string
	Books *model.Book
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
	//res := map[string]string{
	//	th[0]: el[0],
	//	th[1]: el[1],
	//	th[2]: el[2],
	//}
	var res = make(map[string]string, len(tmp))
	for i := 0; i < len(tmp); i++ {
		tmp, _ := url.QueryUnescape(el[i])
		res[th[i]] = tmp
	}
	return res
}

//User新規登録
func (u *DBHandler) NewUser(w http.ResponseWriter, r *http.Request) {
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
func (u *DBHandler) Login(w http.ResponseWriter, r *http.Request) {
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
func (u *DBHandler) SelectPersonalInfo(w http.ResponseWriter, r *http.Request) {
	//return用struct

	var res Res
	body := body(r)

	tmp, _ := strconv.Atoi(body["id"])
	var usr = model.User{
		ID: tmp,
	}

	usrInfo, err := usr.SelectById(u.DB)
	if err != nil {
		fmt.Println("Error!", err)
	}

	booksinfo, err := usr.SelectPersonalBooks(u.DB)
	if err != nil {
		fmt.Println("Error!", err)
	}
	fmt.Println("BOOKSINFO", booksinfo)

	res.ID = usrInfo.ID
	res.Name = usrInfo.Name
	res.Email = usrInfo.Email
	res.Books = booksinfo
	fmt.Println("JKJKJKJKJKJKJKKKKKKKKKKKKK", res)

	marUsr, _ := json.Marshal(res)
	fmt.Fprintf(w, string(marUsr))
}

//本の新規登録。楽天Books API を内部で叩く
func (u *DBHandler) GetBooksInfo(w http.ResponseWriter, r *http.Request) {
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
func (u *DBHandler) RegistBook(w http.ResponseWriter, r *http.Request) {
	body := body(r)

	price, _ := strconv.Atoi(body["price"])
	user_id, _ := strconv.Atoi(body["user_id"])
	var book = model.Book{
		Title:          body["title"],
		Author:         body["author"],
		Price:          price,
		ImgUrl:         body["img_url"],
		RakutenPageUrl: body["rakuten_page_url"],
		User_id:        user_id,
	}
	_, err := book.InsertBook(u.DB)
	if err != nil {
		return
	}
}
