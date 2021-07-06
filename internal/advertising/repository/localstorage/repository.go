package repository

import (
	"errors"
	"github.com/polyanimal/advertising/internal/models"
	"sort"
	"sync"
)

type AdvertisingRepo struct {
	mutex          sync.Mutex
	advertisements map[string]models.Advertisement
}

func NewAdvertisingRepo() *AdvertisingRepo {
	m := make(map[string]models.Advertisement)
	return &AdvertisingRepo{
		advertisements: m,
	}
}

func (r *AdvertisingRepo) GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error) {
	ads := make([]models.Advertisement, 0)

	for _, ad := range r.advertisements {
		ads = append(ads, ad)
	}

	if options.Sort == "by_date" {
		sort.Slice(ads, func(i, j int) bool {
			return ads[i].DateCreate.Before(ads[j].DateCreate)
		})
	} else if options.Sort == "by_price" {
		sort.Slice(ads, func(i, j int) bool {
			return ads[i].Price < ads[j].Price
		})
	} else {
		return nil, errors.New("invalid sort option")
	}

	if options.Order == "descending" {
		for i, j := 0, len(ads)-1; i < j; i, j = i+1, j-1 { //reverse slice
			ads[i], ads[j] = ads[j], ads[i]
		}
	}

	start := (options.PageNumber - 1) * options.ObjectsPerPage
	if start >= len(ads) {
		return nil, errors.New("invalid page")
	}

	end := start + options.ObjectsPerPage
	if end > len(ads) {
		end = len(ads)
	}

	return ads[start:end], nil
}

func (r *AdvertisingRepo) GetAdvertisement(ID string) (models.Advertisement, error) {
	a, exists := r.advertisements[ID]
	if !exists {
		return models.Advertisement{}, errors.New("advertisement not found")
	}
	return a, nil
}

func (r *AdvertisingRepo) CreateAdvertisement(ad models.Advertisement) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.advertisements[ad.ID] = ad

	return nil
}
