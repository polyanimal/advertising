package advertising

import "github.com/polyanimal/advertising/internal/models"

type UseCase interface {
	GetAllAdvertisements() ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) (string, error)
}
