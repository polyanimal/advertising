package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	"github.com/polyanimal/advertising/internal/models"
	"sort"
)

// PgxPoolIface Интерфейс для драйвера БД
type PgxPoolIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
}

type AdvertisingRepo struct {
	db PgxPoolIface
}

func NewAdvertisingRepo(database PgxPoolIface) *AdvertisingRepo {
	return &AdvertisingRepo{
		db: database,
	}
}

func (r *AdvertisingRepo) GetAllAdvertisements(options *models.Options) ([]models.Advertisement, error) {
	ads := make([]models.Advertisement, 0)

	sqlStatement := `
        SELECT * FROM mdb.advertisement
    `
	rows, err := r.db.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ad := models.Advertisement{}
		err = rows.Scan(&ad.ID, &ad.Name, &ad.Description, &ad.PhotoLinks, &ad.Price, &ad.DateCreate)
		if err != nil {
			return nil, err
		}
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
	var ad models.Advertisement

	sqlStatement := `
        SELECT id, name, description, photo_links, price, date_create FROM mdb.advertisement 
        WHERE id=$1
    `
	err := r.db.
		QueryRow(context.Background(), sqlStatement, ID).
		Scan(&ad.ID, &ad.Name, &ad.Description, &ad.PhotoLinks, &ad.Price, &ad.DateCreate)

	if err != nil {
		return models.Advertisement{}, errors.New("advertisement not found")
	}

	return ad, nil
}

func (r *AdvertisingRepo) CreateAdvertisement(ad models.Advertisement) error {
	sqlStatement := `
        INSERT INTO mdb.advertisement (id, name, description, photo_links, price, date_create)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := r.db.
		Exec(context.Background(), sqlStatement, ad.ID, ad.Name, ad.Description, ad.PhotoLinks, ad.Price, ad.DateCreate)

	if err != nil {
		return errors.New("create Advertisement Error")
	}

	return nil
}
