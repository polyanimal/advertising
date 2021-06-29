package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
)

func RegisterHTTPEndpoints(router *gin.Engine, advertisingUC advertising.UseCase) {
	handler := NewHandler(advertisingUC)

	router.GET("/advertisements", handler.GetAllAdvertisements)
	router.GET("/advertisements/:id", handler.GetAdvertisement)
	router.POST("/advertisements/", handler.CreateAdvertisement)
}
