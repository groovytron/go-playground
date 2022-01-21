package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAtd"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name  string             `json:"name"`
	Description  string      `json:"description"`
	Tasks []Task             `json:"-"`
}
