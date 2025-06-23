package main

import (
	"fmt"
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

	taskFromDB, err := taskSrv.Create(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskFromDB)

	// UPDATE
	//t.Name = "updated task"
	//query := "UPDATE tasks SET name=$1 WHERE id=$2"
	//_, err = db.Exec(query, t.Name, t.ID)
	//if err != nil {
	//	panic(err)
	//}

}
