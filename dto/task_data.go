package dto

import (
	"time"
)

type TaskData struct {
	Id          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryId  string     `json:"category_id"`
	Category    string     `json:"category"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TaskRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=500"`
	CategoryID  string `json:"category_id" validate:"required,uuid4"`
}

type IdTaskRequest struct {
	Id string `json:"id_task" validate:"required,uuid4"`
}
