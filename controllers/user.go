package controllers

import (
	"github.com/fmo/tasks/domain/users"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	userSvc *users.Service
}

func NewUser(us *users.Service) *User {
	return &User{us}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
	}
	u.userSvc.Create()

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
