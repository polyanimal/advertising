package advertising

import "github.com/polyanimal/advertising/internal/models"

type Repository interface {
	GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) error
}
