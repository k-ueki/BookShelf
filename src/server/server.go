package main

import (
	"database/sql"
	"fmt"
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
	} else {
		fmt.Println("seccess for connection!!")
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
	uctr := &ctrl.User{DB: db, Stream: userInfo}
	//fmt.Println(uctr)
	r.HandleFunc("/", uctr.Login)
	r.HandleFunc("/signup/", uctr.NewUser)
	r.HandleFunc("/top/", top)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))
}

func top(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "top/index", r)
}
