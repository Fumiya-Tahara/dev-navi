package entity

import "time"

type Product struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Memo     string    `json:"memo"`
}