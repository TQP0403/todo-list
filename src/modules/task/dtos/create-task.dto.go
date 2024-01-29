package dtos

import "TQP0403/todo-list/src/models"

type CreateTaskDto struct {
	ID      int               `json:"-"`
	UserID  int               `json:"-"`
	Title   string            `json:"title,omitempty" binding:"required,max:255"`
	Content string            `json:"content,omitempty" binding:"required,max:255"`
	Status  models.TaskStatus `json:"status,omitempty"`
}
