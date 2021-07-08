package repository

import (
	"context"
	"github.com/pashagolub/pgxmock"
	"github.com/polyanimal/advertising/internal/models"
	"testing"
	"time"
)

func TestAdvertisingRepo(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close(context.Background())
	repo := NewAdvertisingRepo(mockDB)

	//testOptionsStruct := models.Options{
	//	Sort:           "by_date",
	//	Order:          "ascending",
	//	ObjectsPerPage: 10,
	//	PageNumber:     1,
	//}

	testAdd := models.Advertisement{
		ID: "xxxxx",
		Name: "1",
		Description: "car",
		PhotoLinks: []string{"1", "2"},
		Price: 10000,
		DateCreate: time.Now(),
	}

	t.Run("CreateAd", func(t * testing.T){
		mockDB.ExpectExec("INSERT INTO").WithArgs(pgxmock.AnyArg(), testAdd.Name, testAdd.Description,
			testAdd.PhotoLinks, testAdd.Price, pgxmock.AnyArg()).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		if _, err = repo.CreateAdvertisement(testAdd); err != nil {
			t.Errorf("error was not expected while updating stats: %s", err)
		}

		if err := mockDB.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}
