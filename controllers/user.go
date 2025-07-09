package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/add_user.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "cant parse template", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, "cant run the template", http.StatusInternalServerError)
	}
}
