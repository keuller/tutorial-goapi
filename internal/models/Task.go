package models

import "time"

type Task struct {
	ID          string `db:"id"`
	Title       string
	Description string
	Done        bool      `db:"done"`
	CreatedAt   time.Time `db:"created_at"`
}

type TaskResource struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type AddTask struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTask struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
