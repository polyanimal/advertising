package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/polyanimal/advertising/internal/advertising/mocks"
	"github.com/polyanimal/advertising/internal/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoviesUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	uc := NewAdvertisingUC(repo)

	ID := uuid.NewV4().String()

	ad := models.Advertisement{
		Name: "1",
		Description: "car",
		PhotoLinks: []string{"1", "2"},
		Price: 10000,
	}

	t.Run("CreateAd", func(t *testing.T) {
		repo.EXPECT().CreateAdvertisement(ad).Return(ID, nil)
		newID, err := uc.CreateAdvertisement(ad)
		assert.Equal(t, ID, newID)
		assert.NoError(t, err)
	})

}
