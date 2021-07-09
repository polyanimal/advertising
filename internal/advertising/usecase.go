package advertising

import "github.com/polyanimal/advertising/internal/models"


// UseCase go:generate mockgen -destination=mocks/usecase_mock.go -package=mocks . UseCase
type UseCase interface {
	GetAllAdvertisements(options *models.Options) ([]models.AdFeedItem, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) (string, error)
}
