package entity

import (
	"time"
)

type Milestone struct {
	ID       uint      `db:"id"`
	Title    string    `db:"title"`
	Deadline time.Time `db:"deadline"`
	Status   uint      `db:"status"`
}
