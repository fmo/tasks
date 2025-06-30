package main

import (
	"errors"
	"github.com/fmo/tasks/database/connections"
	"github.com/fmo/tasks/domain/tasks"
	"github.com/fmo/tasks/models"
	"github.com/fmo/tasks/project_errors"
	"github.com/joho/godotenv"
	"log"
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
		Owner:    "joe",
		Priority: 1,
	}

	db, err := connections.NewPostgresConnection(os.Getenv("DB_CONN_STRING"))
	if err != nil {
		panic(err)
	}

	taskRepo := tasks.NewRepository(db)
	taskSrv := tasks.NewService(taskRepo)

	taskFromDB, err := taskSrv.Create(&t)
	if err != nil {
		handleError(err)
	}

	taskFromDB.Name = "updated name"

	// Update task
	err = taskSrv.Update(taskFromDB)
	if err != nil {
		panic(err)
	}
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
