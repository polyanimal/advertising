package models

import "time"

type Advertisement struct {
	ID          string       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PhotoLinks  []string  `json:"photo_links"`
	Price       uint      `json:"price"`
	DateCreate  time.Time `json:"date_create"`
}
