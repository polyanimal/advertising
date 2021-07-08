package advertising

import "github.com/polyanimal/advertising/internal/models"


// UseCase go:generate mockgen -destination=mocks/usecase.go -package=mocks . UseCase
type UseCase interface {
	GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) (string, error)
}
