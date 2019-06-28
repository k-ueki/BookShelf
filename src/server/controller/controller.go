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
type Community struct {
	Community_ID   int
	Community_Name string
}

//return用
type Res struct {
	ID     int          `json:id`
	Name   string       `json:name`
	UserId string       `json:userid`
	Email  string       `json:email`
	Books  []model.Book `json:books`
}

//http.Requestからbodyをmapで返す
func body(r *http.Request) map[string]string {
	fmt.Println(r.ContentLength)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
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
		UserId:   body["userid"],
		Password: body["password"],
		Email:    body["email"],
	}
	flag, err, clu := usr.CheckPreDB(u.DB)
	if err != nil {
		fmt.Println(err)
		return
	}
	if flag == false {
		switch clu {
		case "email":
			fmt.Fprintf(w, "このメールアドレスはすでに登録されています")
			return
		case "userid":
			fmt.Fprintf(w, "このユーザーIDはすでに使用されています")
			return
		case "useridemail":
			fmt.Fprintf(w, "このメールアドレスとユーザーIDはすでに使用されています")
			return
		}
	}

	_, err = usr.Insert(u.DB)
	if err != nil {
		fmt.Println("FAILED!!!", err)
		return
	}

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
	res, err := usr.GetInfoByEmailPass(u.DB)
	if err != nil {
		fmt.Println("ERROR!", err)
		fmt.Fprintf(w, "ERROR")
		return
	}
	fmt.Fprintf(w, fmt.Sprint(res.ID))
}

//Personal Info をIDから取得
func (u *DBHandler) SelectPersonalInfo(w http.ResponseWriter, r *http.Request) {
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
	res.UserId = usrInfo.UserId
	res.Name = usrInfo.Name
	res.Email = usrInfo.Email
	res.Books = booksinfo

	marUsr, _ := json.Marshal(res)
	fmt.Fprintf(w, string(marUsr))
}
func (u *DBHandler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	body := body(r)
	fmt.Println("JJ", body)

	tmp, _ := strconv.Atoi(body["delid"])
	var book = model.Book{
		Id: tmp,
	}
	err := book.DeleteBook(u.DB)
	if err != nil {
		fmt.Println("Error! Deleting!", err)
	}
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
func (u *DBHandler) SelectAllPerson(h http.ResponseWriter, r *http.Request) {
	var allperson = model.User{}

	res, _ := allperson.SelectAllPerson(u.DB)
	mar, _ := json.Marshal(res)
	fmt.Fprintf(h, string(mar))
}

//=========================
// Community
// ========================

//func (u *DBHandler) Community(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		u.SelectAllPerson(w, r)
//	case "POST":
//		u.ResisterCommunity(w, r)
//	}
//
//}

//func (u *DBHandler) ResisterCommunity(w http.ResponseWriter, r *http.Request) {
//	body := body(r)
//
//	tmp := len(body["CommunityMembers"])
//	members := strconv.Itoa(tmp)
//
//	var com = model.Community{
//		Name: body["CommunityName"],
//		//Pass:  body["Pass"],
//		Users: members,
//	}
//	if body["Pass"] != "" {
//		com.Pass = body["Pass"]
//	}
//
//	var usr = model.User{
//		Name: body["CommunityMembers"],
//	}
//	userinfo, _ := usr.SelectUserIDByName(u.DB)
//	fmt.Println("userID", userinfo)
//
//	switch members {
//	case "1":
//		com.User1 = body["CommunityMembers"]
//	}
//
//	err := com.Resister(u.DB)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Fprintf(w, "OK")
//}
