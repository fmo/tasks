package main

import (
	"github.com/fmo/tasks/database/connections"
	"github.com/fmo/tasks/domain/tasks"
	"github.com/fmo/tasks/models"
)

import _ "github.com/lib/pq"

func main() {
	t := models.Task{
		Name:     "first task",
		Owner:    "joe",
		Priority: 1,
	}

	db, err := connections.NewPostgresConnection("user=user password=password dbname=tasks port=5433 sslmode=disable")
	if err != nil {
		panic(err)
	}

	taskRepo := tasks.NewRepository(db)
	taskSrv := tasks.NewService(taskRepo)

	taskFromDB, err := taskSrv.Create(&t)
	if err != nil {
		panic(err)
	}

	taskFromDB.Name = "updated name"

	// Update task
	err = taskSrv.Update(taskFromDB)
	if err != nil {
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
