package models

type Options struct {
	Sort           string `json:"sort"`
	Order          string `json:"order"`
	ObjectsPerPage int    `json:"objects_per_page"`
	PageNumber     int    `json:"page"`
}
