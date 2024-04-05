package entity

import "time"

type Milestone struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"project_id"`
	Title     string    `json:"title"`
	Deadline  time.Time `json:"deadline"`
}
