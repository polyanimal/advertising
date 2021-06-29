package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
)

type Handler struct {
	advertisingUC advertising.UseCase
}

func NewHandler(useCase advertising.UseCase) *Handler {
	return &Handler{
		advertisingUC: useCase,
	}
}

type AdvertisementsResponse struct {
	CurrentPage    int                    `json:"current_page"`
	PagesNumber    int                    `json:"pages_number"`
	Advertisements []models.Advertisement `json:"advertisements"`
}

func (h *Handler) GetAllAdvertisements(ctx *gin.Context) {

}

func (h *Handler) GetAdvertisement(ctx *gin.Context) {

}

func (h *Handler) CreateAdvertisement(ctx *gin.Context) {

}
