package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
	"github.com/polyanimal/advertising/pkg/util"
	"net/http"
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

type CreateAdvertisementResponse struct {
	ID string `json:"id"`
}

func (h *Handler) GetAllAdvertisements(ctx *gin.Context) {

}

func (h *Handler) GetAdvertisement(ctx *gin.Context) {

}

func (h *Handler) CreateAdvertisement(ctx *gin.Context) {
	ad := new(models.Advertisement)
	err := ctx.BindJSON(ad)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.advertisingUC.CreateAdvertisement(*ad)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := CreateAdvertisementResponse{
		ID: id,
	}

	ctx.JSON(http.StatusCreated, response)
}
