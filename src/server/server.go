package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, tmpl string, r *http.Request) {
	//	t := template.Must(template.ParseFiles("client/src/" + tmpl + ".js"))
	t := template.Must(template.ParseFiles("templates/" + tmpl + ".html"))
	//t := template.Must(template.ParseFiles("templates/" + tmpl + ".vue"))
	var filename string
	if strings.Contains(tmpl, "/") {
		file := strings.Split(tmpl, "/")
		filename = file[len(file)-1]
	} else {
		filename = tmpl
	}

	fmt.Println(tmpl)
	fmt.Println(filename)
	//if err := t.ExecuteTemplate(w, filename+".vue", "h"); err != nil {
	if err := t.ExecuteTemplate(w, filename+".html", "h"); err != nil {
		//if err := t.ExecuteTemplate(w, filename+".js", "h"); err != nil {
		log.Fatal(err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	//	renderTemplate(w, "index", r)
	//renderTemplate(w, "main", r)
}
func signup(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "signup/index", r)
}
func top(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "top/index", r)
}

func main() {
	var DSN string = "root:@/uchihon"

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	r := mux.NewRouter()
	//handler
	r.HandleFunc("/", index)
	r.HandleFunc("/signup/", signup)
	r.HandleFunc("/top/", top)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.ListenAndServe(":8888", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))

}
