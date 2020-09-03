package main

import (
	"log"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		uId := r.FormValue("id")
		db := conn()
		defer db.Close()
		rows, err := db.Prepare(
			"DELETE FROM users WHERE id=?",
		)
		if err != nil {
			panic(err.Error())
		}
		rows.Exec(uId)
		log.Println("DELETE: id=", uId)
		http.Redirect(w, r, "/", 301)
	} else {
		tmpl.ExecuteTemplate(w, "NotFound", 404)
	}
}
