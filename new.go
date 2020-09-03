package main

import (
	"log"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := conn()
		defer db.Close()
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		ins, err := db.Prepare(
			"INSERT INTO users(`nome`,`email`) VALUES(?,?)",
		)
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(nome, email)
		log.Printf(
			"\nNew User->\nNome:%s, Email: %s",
			nome, email,
		)
		http.Redirect(w, r, "/", 301)
	}
	tmpl.ExecuteTemplate(w, "New", 200)
}
