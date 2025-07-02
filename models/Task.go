package models

import "time"

type Task struct {
	ID        int
	Name      string
	Priority  int
	User      *User
	CreatedAt time.Time
}
