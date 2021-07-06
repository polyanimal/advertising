package usecase

import (
	"errors"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
	constants "github.com/polyanimal/advertising/pkg/const"
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

func (uc *AdvertisingUC) GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error) {
	if options.PageNumber <= 0 {
		return nil, errors.New("invalid page")
	}

	options.ObjectsPerPage = constants.AdsPerPage

	return uc.repository.GetAllAdvertisements(options)
}

func (uc *AdvertisingUC) GetAdvertisement(ID string) (models.Advertisement, error) {
	return uc.repository.GetAdvertisement(ID)
}
func (uc *AdvertisingUC) CreateAdvertisement(ad models.Advertisement) (string, error) {

	if ad.Name == "" || ad.Description == "" || len(ad.PhotoLinks) == 0 {  //TODO ?
		return "", errors.New("invalid advertisement fields")
	}

	ID := uuid.NewV4().String()
	ad.ID = ID
	ad.DateCreate = time.Now()

	err := uc.repository.CreateAdvertisement(ad)

	return ID, err
}
