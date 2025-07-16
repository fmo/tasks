package controllers

import (
	"fmt"
	"github.com/fmo/tasks/domain/users"
	"github.com/fmo/tasks/models"
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

		userFromDB, err := u.userSvc.Create(&models.User{Username: username, Email: email, Password: password})
		if err != nil {
			log.Println(err)
		}

		fmt.Println(userFromDB)
	}

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
