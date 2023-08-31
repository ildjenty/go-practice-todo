package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	Title       string         `json:"title" gorm:"type:varchar(100)"`
	Description string         `json:"description" gorm:"type:text"`
	SortOrder   uint           `json:"sort_order"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}