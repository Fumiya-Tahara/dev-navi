package entity

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Deadline  time.Time      `json:"deadline"`
	Status    uint           `json:"status"`
	Memo      string         `json:"memo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
