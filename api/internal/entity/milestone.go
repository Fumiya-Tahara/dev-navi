package entity

import (
	"time"

	"gorm.io/gorm"
)

type Milestone struct {
	ID        uint           `json:"id"`
	ProjectID uint           `json:"project_id"`
	Title     string         `json:"title"`
	Deadline  time.Time      `json:"deadline"`
	Status    uint           `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
