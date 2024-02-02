package dtos

import "TQP0403/todo-list/src/models"

type UpdateTaskDto struct {
	ID      int               `json:"-"`
	Title   string            `json:"title" binding:"max=255"`
	Content string            `json:"content" binding:"max=255"`
	Status  models.TaskStatus `json:"status,omitempty"`
}
