package advertising

import "github.com/polyanimal/advertising/internal/models"

type Repository interface {
	GetAllAdvertisements() ([]models.Advertisement, error)
	GetAdvertisement() (models.Advertisement, error)
	CreateAdvertisement() error
}
