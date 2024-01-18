package dtos

import "TQP0403/todo-list/src/models"

type UpdateTaskDto struct {
	ID      int               `json:"-"`
	Title   string            `json:"title"`
	Content string            `json:"content"`
	Status  models.TaskStatus `json:"status,omitempty"`
}
