package main

import (
	"errors"
	"github.com/fmo/tasks/controllers"
	"github.com/fmo/tasks/database/connections"
	"github.com/fmo/tasks/domain/tasks"
	"github.com/fmo/tasks/domain/users"
	"github.com/fmo/tasks/project_errors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

import _ "github.com/lib/pq"

func main() {
	if os.Getenv("environment") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	db, err := connections.NewPostgresConnection(os.Getenv("DB_CONN_STRING"))
	if err != nil {
		panic(err)
	}

	taskRepo := tasks.NewRepository(db)
	taskSrv := tasks.NewService(taskRepo)

	userRepo := users.NewRepository(db)
	userSvc := users.NewService(userRepo)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	userController := controllers.NewUser(userSvc)
	taskController := controllers.NewTask(taskSrv)

	http.Handle("/add-user", http.HandlerFunc(userController.Create))
	http.Handle("/add-task", http.HandlerFunc(taskController.Create))

	http.ListenAndServe(":8080", nil)
}

func handleError(err error) {
	if errors.Is(err, project_errors.ErrPriorityRange) {
		log.Printf("%v", err)
	} else {
		panic(err)
	}
}

// GO
// GO
// GO
// GO
// GO
// GO
// GO
// GO
