package models

import (
	"TQP0403/todo-list/src/helper"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// enum TaskStatus
type TaskStatus int8

const (
	InProccess TaskStatus = iota
	Done
	Cancel
)

// enum TaskStatus implement stringer
var taskStatusStr = []string{"InProccess", "Done", "Cancel"}

func (e TaskStatus) String() string {
	return taskStatusStr[e]
}

type Task struct {
	ID        int            `json:"id" gorm:"primarykey"`
	UserID    int            `json:"userId" gorm:"column:user_id"`
	User      User           `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title     string         `json:"title" gorm:"column:title;size:255"`
	Content   string         `json:"content" gorm:"column:content"`
	Status    TaskStatus     `json:"status" gorm:"column:status"`
	CreatedAt *time.Time     `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index"`
}

func (Task) TableName() string {
	return fmt.Sprintf("%s.%s", helper.GetDefaultEnv("DB_SCHEMA", "public"), "tasks")
}

func (task Task) String() string {
	return fmt.Sprintf(
		"[task] id:%d - user_id:%d - title:%s - status:%s - content:%s\n",
		task.ID,
		task.UserID,
		task.Title,
		task.Status,
		task.Content,
	)
}

func (task *Task) MarshalBinary() (data []byte, err error) {
	return json.Marshal(task)
}

func (task *Task) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, task)
}

func (task *Task) Clone() *Task {
	return &Task{
		ID:        task.ID,
		UserID:    task.UserID,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeletedAt,
	}
}
