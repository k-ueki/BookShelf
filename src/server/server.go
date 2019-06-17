package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	ctrl "github.com/k-ueki/app2/src/server/controller"
	"github.com/k-ueki/app2/src/server/model"
)

func main() {
	var DSN string = "root:@/uchihon"

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(db)
	defer db.Close()

	//CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	r := mux.NewRouter()
	//handler
	userInfo := make(chan *model.User)
	uctr := &ctrl.DBHandler{DB: db, Stream: userInfo}
	//fmt.Println(uctr)
	r.HandleFunc("/", uctr.Login)
	r.HandleFunc("/signup/", uctr.NewUser)
	r.HandleFunc("/top/", uctr.SelectPersonalInfo)
	r.HandleFunc("/top/del/", uctr.DeleteBookByID)
	r.HandleFunc("/top/bookapi/", uctr.GetBooksInfo)
	//r.HandleFunc("/top/booksInfo/", uctr.DispBooksDetail)
	r.HandleFunc("/regist/book/", uctr.RegistBook)

	//r.HandleFunc("/community/add/", uctr.SelectAllPerson)
	r.HandleFunc("/community/add/", uctr.Community)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))
}
