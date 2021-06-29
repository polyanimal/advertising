package usecase

import (
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
)

type AdvertisingUC struct {
	repository advertising.Repository
}

func NewAdvertisingUC(repository advertising.Repository) *AdvertisingUC {
	return &AdvertisingUC{
		repository: repository,
	}
}

func (uc *AdvertisingUC) GetAllAdvertisements() ([]models.Advertisement, error) {
	return nil, nil
}

func (uc *AdvertisingUC) GetAdvertisement() (models.Advertisement, error) {
	return models.Advertisement{}, nil
}
func (uc *AdvertisingUC) CreateAdvertisement() error {
	return nil
}
