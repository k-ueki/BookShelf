package server

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	controller "github.com/k-ueki/app2/src/server/controller"
)

type Server struct {
	// db     *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	fmt.Println("hoge")
	s.router = s.Route()
}

func (s *Server) Route() *mux.Router {
	var DSN string = "root:@/uchihon"
	r := mux.NewRouter()

	db, err := sqlx.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	//handler
	// userInfo := make(chan *model.User)

	Controller := &controller.DBHandler{DB: db}

	r.Methods(http.MethodGet).Path("/top/{uid}").Handler(AppHandler{Controller.Index})

	r.Methods(http.MethodGet).Path("/books/{title}").Handler(AppHandler{Controller.GetBooks})
	r.Methods(http.MethodPost).Path("/books").Handler(AppHandler{Controller.PostBooks})

	r.Methods(http.MethodGet).Path("/community/{uid}").Handler(AppHandler{Controller.GetCommunities})
	r.Methods(http.MethodPost).Path("/community").Handler(AppHandler{Controller.PostCommunities})

	// r.HandleFunc("/signup/", uctr.NewUser)
	// r.HandleFunc("/top/", uctr.SelectPersonalInfo)
	// r.HandleFunc("/top/del/", uctr.DeleteBookByID)
	// //r.HandleFunc("/top/booksInfo/", uctr.DispBooksDetail)
	// r.HandleFunc("/community/add/", uctr.SelectAllPerson)
	//r.HandleFunc("/community/add/", uctr.Community)

	// r.HandleFunc("/", uctr.Login)
	// r.HandleFunc("/regist/book/", uctr.RegistBook)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))

	return r
}
