package entity

import (
	"time"
)

type Project struct {
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Status   uint      `json:"status"`
	Memo     string    `json:"memo"`
}
