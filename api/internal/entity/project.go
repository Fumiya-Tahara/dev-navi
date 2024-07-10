package entity

import (
	"time"
)

type Project struct {
	ID         uint        `gorm:"id"`
	Title      string      `gorm:"title"`
	Deadline   time.Time   `gorm:"deadline"`
	Status     uint        `gorm:"status"`
	Memo       string      `gorm:"memo"`
	CreatedAt  time.Time   `gorm:"created_at"`
	UpdatedAt  time.Time   `gorm:"updated_at"`
	Users      []User      `gorm:"many2many:user_projects"`
	Milestones []Milestone `gorm:"foreignKey:ProjectID"`
}
