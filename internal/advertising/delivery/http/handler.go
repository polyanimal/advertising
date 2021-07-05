package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/models"
	"github.com/polyanimal/advertising/pkg/util"
	"net/http"
)

type Handler struct {
	useCase advertising.UseCase
}

func NewHandler(useCase advertising.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type AdvertisementsResponse struct {
	CurrentPage    int                    `json:"current_page"`
	PagesNumber    int                    `json:"pages_number"`
	Advertisements []models.Advertisement `json:"advertisements"`
}

type AdvertisementResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Price       uint     `json:"price"`
	MainPhoto   string   `json:"main_photo"`
	Description string   `json:"description"`
	AllPhotos   []string `json:"all_photos"`
}

type CreateAdvertisementResponse struct {
	ID string `json:"id"`
}

type fieldsParam struct {
	Fields []string `json:"fields"`
}

func (h *Handler) GetAllAdvertisements(ctx *gin.Context) {

}

func (h *Handler) GetAdvertisement(ctx *gin.Context) {
	ID := ctx.Param("id")
	fields := new(fieldsParam)
	errFields := ctx.BindJSON(fields)

	ad, err := h.useCase.GetAdvertisement(ID)
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response := AdvertisementResponse{
		ID:        ad.ID,
		Name:      ad.Name,
		Price:     ad.Price,
		MainPhoto: ad.PhotoLinks[0],
	}

	if errFields == nil {
		for _, s := range fields.Fields {
			switch {
			case s == "Description":
				response.Description = ad.Description
			case s == "AllPhotos":
				response.AllPhotos = ad.PhotoLinks
			}
		}
	}

}

func (h *Handler) CreateAdvertisement(ctx *gin.Context) {
	ad := new(models.Advertisement)
	err := ctx.BindJSON(ad)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.useCase.CreateAdvertisement(*ad)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := CreateAdvertisementResponse{
		ID: id,
	}

	ctx.JSON(http.StatusCreated, response)
}