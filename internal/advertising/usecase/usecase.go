package usecase

import (
	"errors"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
	constants "github.com/polyanimal/advertising/pkg/const"
)

type AdvertisingUC struct {
	repository advertising.Repository
}

func NewAdvertisingUC(repository advertising.Repository) *AdvertisingUC {
	return &AdvertisingUC{
		repository: repository,
	}
}

func (uc *AdvertisingUC) GetAllAdvertisements(options *models.Options) ([]models.AdFeedItem, error) {
	if options.PageNumber <= 0 {
		return nil, errors.New("invalid page")
	}

	options.ObjectsPerPage = constants.AdsPerPage

	ads, err := uc.repository.GetAllAdvertisements(options)
	if err != nil {
		return []models.AdFeedItem{}, err
	}

	adFeed := make([]models.AdFeedItem, 0)
	for _, ad := range ads {
		adFeed = append(adFeed, models.AdFeedItem{Name: ad.Name, MainPhoto: ad.PhotoLinks[0], Price: ad.Price})
	}

	return adFeed, nil
}

func (uc *AdvertisingUC) GetAdvertisement(ID string) (models.Advertisement, error) {
	return uc.repository.GetAdvertisement(ID)
}
func (uc *AdvertisingUC) CreateAdvertisement(ad models.Advertisement) (string, error) {

	if ad.Name == "" || ad.Description == "" || len(ad.PhotoLinks) == 0 {  //TODO ?
		return "", errors.New("invalid advertisement fields")
	}

	ID, err := uc.repository.CreateAdvertisement(ad)
	return ID, err
}
