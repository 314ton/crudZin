package main

import (
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	db := conn()
	defer db.Close()

	id := r.URL.Query().Get("id")
	user := users{}

	if id != "" {
		resp, err := db.Query(
			"SELECT * FROM users where id=?", id,
		)
		if err != nil {
			panic(err.Error)
		}
		for resp.Next() {
			var id int
			var nome, email string
			err := resp.Scan(&id, &nome, &email)
			if err != nil {
				panic(err.Error())
			}
			user.Id = id
			user.Nome = nome
			user.Email = email
		}
	}
	if r.Method == "POST" {
		id := r.FormValue("uId")
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		up, err := db.Prepare(
			"UPDATE users set nome=?, email=? WHERE id=? ",
		)
		if err != nil {
			panic(err.Error())
		}
		up.Exec(nome, email, id)
		log.Printf(
			"\nUpdate-> nome:%s | email: %s\n",
			nome, email,
		)
		http.Redirect(w, r, "/", 301)
	}
	tmpl.ExecuteTemplate(w, "Update", user)
}
