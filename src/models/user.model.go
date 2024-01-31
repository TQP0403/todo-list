package models

import (
	"TQP0403/todo-list/src/helper"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int            `json:"id" gorm:"primarykey"`
	DisplayName string         `json:"displayName" gorm:"column:display_name;size:255"`
	Username    string         `json:"username" gorm:"column:username;size:255;index:unique"`
	Password    string         `json:"-" gorm:"column:password;size:255"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return fmt.Sprintf("%s.%s", helper.GetDefaultEnv("DB_SCHEMA", "public"), "users")
}

func (user User) String() string {
	return fmt.Sprintf(
		"[user] id:%d - display_name:%s - username:%s\n",
		user.ID,
		user.DisplayName,
		user.Username,
	)
}

func (user *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(user)
}

func (user *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, user)
}

func (user *User) Clone() *User {
	return &User{
		ID:          user.ID,
		DisplayName: user.DisplayName,
		Username:    user.Username,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
}
