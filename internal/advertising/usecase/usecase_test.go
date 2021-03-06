package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/polyanimal/advertising/internal/advertising/mocks"
	"github.com/polyanimal/advertising/internal/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvertisingUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	uc := NewAdvertisingUC(repo)

	ID := uuid.NewV4().String()

	testErr := errors.New("test error")

	testAdd := models.Advertisement{
		Name: "1",
		Description: "car",
		PhotoLinks: []string{"1", "2"},
		Price: 10000,
	}

	invalidAd := models.Advertisement{
		Name: "",
		Description: "",
		PhotoLinks: []string{},
		Price: 0,
	}

	testOptionsStruct := models.Options{
		Sort:           "by_date",
		Order:          "ascending",
		ObjectsPerPage: 10,
		PageNumber:     1,
	}

	t.Run("CreateAd", func(t *testing.T) {
		repo.EXPECT().CreateAdvertisement(testAdd).Return(ID, nil)
		newID, err := uc.CreateAdvertisement(testAdd)
		assert.Equal(t, ID, newID)
		assert.NoError(t, err)
	})

	t.Run("CreateAd - fail", func(t *testing.T) {
		repo.EXPECT().CreateAdvertisement(testAdd).Return(ID, testErr)
		newID, err := uc.CreateAdvertisement(testAdd)
		assert.Equal(t, ID, newID)
		assert.Error(t, err)
	})

	t.Run("CreateAd - invalid", func(t *testing.T) {
		newID, err := uc.CreateAdvertisement(invalidAd)
		assert.Equal(t, "", newID)
		assert.Error(t, err)
	})

	t.Run("GetAd", func(t *testing.T) {
		repo.EXPECT().GetAdvertisement(ID).Return(testAdd, nil)
		ad, err := uc.GetAdvertisement(ID)
		assert.Equal(t, testAdd, ad)
		assert.NoError(t, err)
	})

	t.Run("GetAd - fail", func(t *testing.T) {
		repo.EXPECT().GetAdvertisement(ID).Return(testAdd, testErr)
		ad, err := uc.GetAdvertisement(ID)
		assert.Equal(t, testAdd, ad)
		assert.Error(t, err)
	})

	t.Run("GetAllAds", func(t *testing.T) {
		repo.EXPECT().GetAllAdvertisements(&testOptionsStruct).Return([]models.Advertisement{}, nil)
		ads, err := uc.GetAllAdvertisements(&testOptionsStruct)
		assert.Equal(t, []models.AdFeedItem{}, ads)
		assert.NoError(t, err)
	})

	t.Run("GetAllAds", func(t *testing.T) {
		_, err := uc.GetAllAdvertisements(&models.Options{PageNumber: -1})
		assert.Error(t, err)
	})

}
