package advertising

import "github.com/polyanimal/advertising/internal/models"

type UseCase interface {
	GetAllAdvertisements() ([]models.Advertisement, error)
	GetAdvertisement() (models.Advertisement, error)
	CreateAdvertisement() error
}
