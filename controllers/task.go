package controllers

import (
	"github.com/fmo/tasks/domain/tasks"
	"net/http"
)

type Task struct {
	taskSvc *tasks.Service
}

func NewTask(ts *tasks.Service) *Task {
	return &Task{ts}
}

func (t *Task) Create(w http.ResponseWriter, r *http.Request) {

}
