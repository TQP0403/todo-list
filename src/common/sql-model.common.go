package common

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time     `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index"`
}

type IModel interface {
	// print model
	String() string
	// gorm table name
	TableName() string
}

type IJsonModel interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
	// to json string
	Marshal() (string, error)
	// fromjson string
	Unmarshal(jsonStr string) error
}
