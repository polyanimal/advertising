package advertising

import "github.com/polyanimal/advertising/internal/models"

type UseCase interface {
	GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) (string, error)
}
