package models

import (
	"TQP0403/todo-list/src/helper"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TaskStatus int8

const (
	TaskStatusInProccess TaskStatus = iota
	TaskStatusDone
	TaskStatusPause
)

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
		"[task] id:%d - user_id:%d - title:%s - status:%d - content:%s\n",
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

func (task *Task) Marshal() (string, error) {
	// struct to string
	if data, err := json.Marshal(task); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (task *Task) Unmarshal(jsonStr string) error {
	// string to struct
	if err := json.Unmarshal([]byte(jsonStr), &task); err != nil {
		return err
	}
	return nil
}
