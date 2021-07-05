package models

import "time"

type Advertisement struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PhotoLinks  []string  `json:"photo_links"`
	Price       uint      `json:"price"`
	DateCreate  time.Time `json:"date_create"`
}
