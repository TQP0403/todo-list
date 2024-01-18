package dtos

import "TQP0403/todo-list/src/models"

type CreateTaskDto struct {
	ID      int               `json:"-"`
	UserID  int               `json:"-"`
	Title   string            `json:"title,omitempty"`
	Content string            `json:"content,omitempty"`
	Status  models.TaskStatus `json:"status,omitempty"`
}
