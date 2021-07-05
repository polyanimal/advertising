package repository

import (
	"errors"
	"github.com/polyanimal/advertising/internal/models"
	"sync"
)

type AdvertisingRepo struct {
	mutex          sync.Mutex
	advertisements map[string]models.Advertisement
}

func NewAdvertisingRepo() *AdvertisingRepo {
	m := make(map[string]models.Advertisement)
	return &AdvertisingRepo{
		advertisements: m,
	}
}

func (r *AdvertisingRepo) GetAllAdvertisements() ([]models.Advertisement, error) {
	return nil, nil
}

func (r *AdvertisingRepo) GetAdvertisement(ID string) (models.Advertisement, error) {
	a, exists := r.advertisements[ID]
	if !exists {
		return models.Advertisement{}, errors.New("advertisement not found")
	}
	return a, nil
}

func (r *AdvertisingRepo) CreateAdvertisement(ad models.Advertisement) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.advertisements[ad.ID] = ad

	return nil
}
