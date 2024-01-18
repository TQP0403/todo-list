package models

import (
	"TQP0403/todo-list/src/config"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int            `json:"id" gorm:"primarykey"`
	DisplayName string         `json:"displayName" gorm:"column:display_name;size:255"`
	Username    string         `json:"username" gorm:"column:username;size:255;index:unique"`
	Password    string         `json:"-" gorm:"column:password"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return fmt.Sprintf("%s.%s", config.Getenv("DB_SCHEMA", "public"), "users")
}

func (user User) String() string {
	return fmt.Sprintf(
		"[task] id:%d - display_name:%s - username:%s\n",
		user.ID,
		user.DisplayName,
		user.Username,
	)
}

func (user User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

func (user User) Marshal() (string, error) {
	// struct to string
	if data, err := json.Marshal(user); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (user *User) Unmarshal(jsonStr string) error {
	// string to struct
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		return err
	}
	return nil
}
