package entity

import (
	"time"
)

type Milestone struct {
	ID        uint      `gorm:"id"`
	Title     string    `gorm:"title"`
	Deadline  time.Time `gorm:"deadline"`
	Status    uint      `gorm:"status"`
	ProjectID uint      `gorm:"project_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
