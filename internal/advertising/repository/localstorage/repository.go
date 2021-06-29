package repository

import (
	"github.com/polyanimal/advertising/internal/models"
	"sync"
)

type AdvertisingRepo struct {
	mu sync.Mutex
	advertisements map[int]models.Advertisement
}

func NewAdvertisingRepo() *AdvertisingRepo {
	m := make(map[int]models.Advertisement)
	return &AdvertisingRepo{
		advertisements: m,
	}
}