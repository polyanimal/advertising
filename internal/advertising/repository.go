package advertising

import "github.com/polyanimal/advertising/internal/models"

// Repository go:generate mockgen -destination=mocks/repository.go -package=mocks . Repository
type Repository interface {
	GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error)
	GetAdvertisement(ID string) (models.Advertisement, error)
	CreateAdvertisement(models.Advertisement) error
}
