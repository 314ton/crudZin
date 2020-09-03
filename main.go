package main

import (
	"database/sql"
	//_ "github.com/314ton/crudZin/update"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

type users struct {
	Id    int
	Nome  string
	Email string
}

func conn() (db *sql.DB) {
	dbDriver := "mysql"
	user := "go"
	passwd := "golang"
	dbname := "go"
	db, err := sql.Open(
		dbDriver,
		user+":"+passwd+"@/"+dbname,
	)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	static := http.FileServer(http.Dir("./static"))
	http.Handle(
		"/static/",
		http.StripPrefix("/static/",
			static),
	)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/new", New)
	http.HandleFunc("/", Index)
	log.Println("Servidor foi iniciado!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Rodando na porta 8080!")
}
