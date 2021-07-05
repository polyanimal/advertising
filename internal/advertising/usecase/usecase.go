package usecase

import (
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
	"github.com/satori/go.uuid"
	"time"
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

func (uc *AdvertisingUC) GetAdvertisement(ID string) (models.Advertisement, error) {
	return uc.repository.GetAdvertisement(ID)
}
func (uc *AdvertisingUC) CreateAdvertisement(ad models.Advertisement) (string, error) {
	ID := uuid.NewV4().String()
	ad.ID = ID
	ad.DateCreate = time.Now()

	err := uc.repository.CreateAdvertisement(ad)

	return ID, err
}
