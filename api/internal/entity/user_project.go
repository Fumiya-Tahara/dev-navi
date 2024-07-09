package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserProject struct {
	UserID    uint           `json:"user_id"`
	ProjectID uint           `json:"project_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
