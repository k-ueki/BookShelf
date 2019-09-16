package controller_user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/repository"
	"github.com/k-ueki/app2/src/server/service"
)

type DBHandler struct {
	DB *sqlx.DB
	// Stream chan *model.User
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

//User新規登録
// func (u *DBHandler) NewUser(w http.ResponseWriter, r *http.Request) {
// 	body := body(r)
// 	var usr = model.User{
// 		Name:     body["name"],
// 		UserId:   body["userid"],
// 		Password: body["password"],
// 		Email:    body["email"],
// 	}
// 	flag, err, clu := usr.CheckPreDB(u.DB)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if flag == false {
// 		switch clu {
// 		case "email":
// 			fmt.Fprintf(w, "このメールアドレスはすでに登録されています")
// 			return
// 		case "userid":
// 			fmt.Fprintf(w, "このユーザーIDはすでに使用されています")
// 			return
// 		case "useridemail":
// 			fmt.Fprintf(w, "このメールアドレスとユーザーIDはすでに使用されています")
// 			return
// 		}
// 	}
//
// 	_, err = usr.Insert(u.DB)
// 	if err != nil {
// 		fmt.Println("FAILED!!!", err)
// 		return
// 	}
//
// 	res, err := usr.SelectIdByNameANDPass(u.DB)
// 	if err != nil {
// 		fmt.Println("ERROR!", err)
// 		return
// 	}
// 	mar, _ := json.Marshal(res)
// 	fmt.Fprintf(w, string(mar))
// }

//ログイン
// func (u *DBHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	body := body(r)
// 	var usr = model.User{
// 		Password: body["pass"],
// 		Email:    body["email"],
// 	}
// 	res, err := usr.GetInfoByEmailPass(u.DB)
// 	if err != nil {
// 		fmt.Println("ERROR!", err)
// 		fmt.Fprintf(w, "ERROR")
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprint(res.ID))
// }

func (u *DBHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	userService := service.NewUserService(u.DB)
	resp, err := userService.Index(uid)
	if err != nil {
		fmt.Println("errorだ!")
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, resp, nil
}

//Personal Info をIDから取得
// func (u *DBHandler) SelectPersonalInfo(w http.ResponseWriter, r *http.Request) {
// 	var res Res
// 	body := body(r)
//
// 	tmp, _ := strconv.Atoi(body["id"])
// 	var usr = model.User{
// 		ID: tmp,
// 	}
//
// 	usrInfo, err := usr.SelectById(u.DB)
// 	if err != nil {
// 		fmt.Println("Error!", err)
// 	}
//
// 	booksinfo, err := usr.SelectPersonalBooks(u.DB)
// 	if err != nil {
// 		fmt.Println("Error!", err)
// 	}
// 	fmt.Println("BOOKSINFO", booksinfo)
//
// 	res.ID = usrInfo.ID
// 	res.UserId = usrInfo.UserId
// 	res.Name = usrInfo.Name
// 	res.Email = usrInfo.Email
// 	res.Books = booksinfo
//
// 	marUsr, _ := json.Marshal(res)
// 	fmt.Fprintf(w, string(marUsr))
// }
//
// func (u *DBHandler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
// 	body := body(r)
// 	fmt.Println("JJ", body)
//
// 	tmp, _ := strconv.Atoi(body["delid"])
// 	var book = model.Book{
// 		Id: tmp,
// 	}
// 	err := book.DeleteBook(u.DB)
// 	if err != nil {
// 		fmt.Println("Error! Deleting!", err)
// 	}
// }
//

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//本の新規登録。楽天Books API を内部で叩く
func (u *DBHandler) GetBooks(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	title := vars["title"]

	Env_load()
	applicationId := os.Getenv("APPLICATION_ID")
	url := "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?"
	url += "applicationId=" + applicationId
	url += "&title=" + title

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	items := model.Resp{}
	books := model.Books{}
	bytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &items)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, nil, err
	}

	books.Books = items.Items

	return http.StatusOK, books, nil
}

func (u *DBHandler) PostBooks(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.PostBookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, nil, err
	}

	var book = &model.Book{
		Title:   reqParam.Title,
		Author:  reqParam.Author,
		Price:   reqParam.Price,
		ImgUrl:  reqParam.ImgUrl,
		PageUrl: &reqParam.PageUrl,
		// User_id:        user_id,
	}
	result, err := repository.Insert(u.DB, *book)
	if err != nil {
		fmt.Println(err)
		// return errInternalServerError
		return http.StatusInternalServerError, nil, err
	}

	bid, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		// return err
		return http.StatusInternalServerError, nil, err
	}
	_, err = repository.RegisterBookAndUser(u.DB, reqParam.UserId, bid)
	if err != nil {
		fmt.Println(err)
		// return err
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil

}

func (u *DBHandler) GetCommunities(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	communityService := service.NewCommunityService(u.DB)
	resp, err := communityService.GetCommunitiesByUid(uid)
	if err != nil {
		fmt.Println("errorだ!")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, resp, nil
}

func (u *DBHandler) PostCommunities(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.PostCommunityRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, nil, err
	}
	uids := reqParam.UserIds
	cid := reqParam.CommunityId

	communityService := service.NewCommunityService(u.DB)
	err := communityService.Create(uids, cid)
	if err != nil {
		fmt.Println("errorだ!")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

// func (u *DBHandler) SelectAllPerson(h http.ResponseWriter, r *http.Request) {
// 	var allperson = model.User{}
//
// 	res, _ := allperson.SelectAllPerson(u.DB)
// 	mar, _ := json.Marshal(res)
// 	fmt.Fprintf(h, string(mar))
