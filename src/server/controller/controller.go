package controller_user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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

//return用
type Res struct {
	ID     int          `json:id`
	Name   string       `json:name`
	UserId string       `json:userid`
	Email  string       `json:email`
	Books  []model.Book `json:books`
}

func (u *DBHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fmt.Println(uid)

	userService := service.NewUserService(u.DB)
	resp, err := userService.Index(uid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, resp, nil
}

func (u *DBHandler) Test(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "test",
		Value: "test,",
	}
	http.SetCookie(w, &cookie)

	// return http.StatusInternalServerError, nil, nil
}

func (u *DBHandler) DiscriminateExists(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	userService := service.NewUserService(u.DB)
	tmp, isExists := userService.IsExists(uid)
	if tmp == nil {
		return http.StatusBadRequest, nil, errors.New("unknown user")
	}
	user_id := strconv.FormatInt(*tmp, 10)
	fmt.Println(user_id)

	// cookie := http.Cookie{
	// 	Name:  "user_id",
	// 	Value: user_id,
	// }
	// http.SetCookie(w, &cookie)

	return http.StatusOK, isExists, nil
}

func (u *DBHandler) RegisterUser(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	req := model.UserReq{
		Uid:  r.FormValue("uid"),
		Name: r.FormValue("name"),
	}
	_, err := repository.Register(u.DB, req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (u *DBHandler) RegisterUserId(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	type user struct {
		UserId   int    `json:"user_id"`
		DispName string `json:"disp_name"`
	}

	req := &user{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	UserService := service.NewUserService(u.DB)
	flagOK, err := UserService.IsOKUserId(req.UserId, req.DispName)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if !flagOK {
		return http.StatusBadRequest, nil, errors.New("already exists")
	}

	if err := UserService.RegisterUserId(req.UserId, req.DispName); err != nil {
		return http.StatusInternalServerError, nil, nil
	}

	return http.StatusOK, nil, nil
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

func Env_load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//本の新規登録。楽天Books API を内部で叩く
func (u *DBHandler) GetBooks(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	title := vars["title"]

	Env_load(os.ExpandEnv("$GOPATH/src/github.com/k-ueki/app2/.env"))
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
		return http.StatusInternalServerError, nil, err
	}

	bid, err := result.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	_, err = repository.RegisterBookAndUser(u.DB, reqParam.UserId, bid)
	if err != nil {
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

func (u *DBHandler) RegisterBook(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.PostBookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var book = &model.Book{
		Title:   reqParam.Title,
		Author:  reqParam.Author,
		Price:   reqParam.Price,
		ImgUrl:  reqParam.ImgUrl,
		PageUrl: &reqParam.PageUrl,
		Isbn:    reqParam.Isbn,
	}

	bookService := service.NewBookService(u.DB)

	bookId, err := bookService.IsExists(book.Isbn)
	if err != true {
		result, err := repository.Insert(u.DB, *book)
		if err != nil {
			return http.StatusInternalServerError, nil, errors.New("failed to register the book")
		}

		tmpId, err := result.LastInsertId()
		if err != nil {
			return http.StatusInternalServerError, nil, errors.New("failed to register the book")
		}

		*bookId = tmpId
	}

	// cookie, _ := r.Cookie("user_id")
	// if cookie == nil {
	// 	return http.StatusBadRequest, nil, errors.New("not login")
	// }
	// uid := cookie.Value

	//暫定
	uid := "6"

	user_id, _ := strconv.ParseInt(uid, 10, 64)
	_, erro := repository.RegisterBookAndUser(u.DB, user_id, *bookId)
	if erro != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to register the book")
	}

	return http.StatusOK, nil, nil
}
