package entity

import "time"

type User struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Projects  []Project `db:"many2many:user_projects"`
}
