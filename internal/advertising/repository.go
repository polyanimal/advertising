package advertising

import "github.com/polyanimal/advertising/internal/models"

type Repository interface {
	GetAllAdvertisements() ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) error
}
