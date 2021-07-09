package models

type AdFeedItem struct {
	Name      string `json:"name"`
	MainPhoto string `json:"photo"`
	Price     uint   `json:"price"`
}
