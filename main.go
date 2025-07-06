package main

import (
	"errors"
	"fmt"
	"github.com/fmo/tasks/controllers"
	"github.com/fmo/tasks/database/connections"
	"github.com/fmo/tasks/models"
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
	t := models.Task{
		Name:     "first task",
		Priority: 1,
	}

	fmt.Println(t)

	db, err := connections.NewPostgresConnection(os.Getenv("DB_CONN_STRING"))
	if err != nil {
		panic(err)
	}

	db.Ping()

	//taskRepo := tasks.NewRepository(db)
	//taskSrv := tasks.NewService(taskRepo)
	//
	//taskFromDB, err := taskSrv.Create(&t)
	//if err != nil {
	//	handleError(err)
	//}
	//
	//taskFromDB.Name = "updated name"

	// Update task
	//err = taskSrv.Update(taskFromDB)
	//if err != nil {
	//	panic(err)
	//}

	userController := controllers.User{}
	http.Handle("/users", http.HandlerFunc(userController.Create))

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
