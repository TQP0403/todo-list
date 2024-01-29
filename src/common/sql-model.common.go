package common

import (
	"encoding"
	"fmt"
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
	// implement stringer: print model
	String() string
	// implement gorm model: table name
	TableName() string

	// json handle
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func IsModel(models ...interface{}) error {
	for _, val := range models {
		if _, ok := val.(IModel); !ok {
			return fmt.Errorf("not a model: %s", val)
		}
	}
	return nil
}
