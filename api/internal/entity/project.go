package entity

import (
	"time"
)

type Project struct {
	ID       uint      `db:"id"`
	Title    string    `db:"title"`
	Deadline time.Time `db:"deadline"`
	Status   uint      `db:"status"`
	Memo     string    `db:"memo"`
}
