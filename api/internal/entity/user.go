package entity

import "time"

type User struct {
	ID        uint      `gorm:"id"`
	Name      string    `gorm:"name"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	Projects  []Project `gorm:"many2many:user_projects"`
}
