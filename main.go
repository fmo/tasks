package main

import (
	"database/sql"
	"fmt"
)

import _ "github.com/lib/pq"

type Task struct {
	id       int
	name     string
	owner    string
	priority int
}

func main() {
	t := Task{
		name:     "first task",
		owner:    "joe",
		priority: 1,
	}

	connStr := "user=user password=password dbname=tasks port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully")

	// CREATE
	query := `
		INSERT INTO tasks (name, owner, priority)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	// UPDATE
	err = db.QueryRow(query, t.name, t.owner, t.priority).Scan(&t.id)
	if err != nil {
		panic(err)
	}
	t.name = "updated task"
	query = "UPDATE tasks SET name=$1 WHERE id=$2"
	_, err = db.Exec(query, t.name, t.id)
	if err != nil {
		panic(err)
	}

}
