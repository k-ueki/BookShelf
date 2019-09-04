package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	controller "github.com/k-ueki/app2/src/server/controller"
)

func main() {
	var DSN string = "root:@/uchihon"
	r := mux.NewRouter()

	db, err := sqlx.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	test := os.Getenv("TEST")
	fmt.Println("test", test)

	//CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	//handler
	// userInfo := make(chan *model.User)

	Controller := &controller.DBHandler{DB: db}
	r.Methods(http.MethodGet).Path("/top/{uid}").HandlerFunc(Controller.Index)

	r.Methods(http.MethodGet).Path("/books/{title}").HandlerFunc(Controller.Get)
	// r.HandleFunc("/", uctr.Login)
	// r.HandleFunc("/signup/", uctr.NewUser)
	// r.HandleFunc("/top/", uctr.SelectPersonalInfo)
	// r.HandleFunc("/top/del/", uctr.DeleteBookByID)
	// //r.HandleFunc("/top/booksInfo/", uctr.DispBooksDetail)
	// r.HandleFunc("/regist/book/", uctr.RegistBook)
	//r.HandleFunc("/community/add/", uctr.SelectAllPerson)
	//r.HandleFunc("/community/add/", uctr.Community)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))
}
