package entity

import "time"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Memo     string    `json:"memo"`
}

type Milestone struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"project_id"`
	Title     string    `json:"title"`
	Deadline  time.Time `json:"deadline"`
}
