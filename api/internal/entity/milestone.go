package entity

import (
	"time"
)

type Milestone struct {
	ID        uint      `db:"id"`
	Title     string    `db:"title"`
	Deadline  time.Time `db:"deadline"`
	Status    uint      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
