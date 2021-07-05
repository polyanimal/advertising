package models

import "time"

type Advertisement struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Publisher   string    `json:"publisher"`
	PhotoLinks  []string  `json:"photo_links"`
	Price       uint      `json:"price"`
	DateCreate  time.Time `json:"date_create"`
}
