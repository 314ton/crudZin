package main

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db := conn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	var usr users
	var usrs []users
	for rows.Next() {
		var id int
		var nome, email string
		err := rows.Scan(&id, &nome, &email)
		if err != nil {
			panic(err.Error())
		}
		usr.Id, usr.Nome, usr.Email = id, nome, email
		usrs = append(usrs, usr)
	}
	log.Println("Exec. Function Index")
	tmpl.ExecuteTemplate(w, "Index", usrs)
}
