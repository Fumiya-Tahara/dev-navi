package entity

import (
	"time"
)

type Milestone struct {
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Status   uint      `json:"status"`
}
