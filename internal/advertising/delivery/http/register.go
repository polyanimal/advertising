package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/advertising"
	"github.com/polyanimal/advertising/internal/middleware"
)

func RegisterHTTPEndpoints(router *gin.Engine, advertisingUC advertising.UseCase, validationMiddleware middleware.Validation) {
	handler := NewHandler(advertisingUC)

	router.POST("/advertisements/", validationMiddleware.CheckValid(), handler.CreateAdvertisement)
	router.GET("/advertisements", handler.GetAllAdvertisements)
	router.GET("/advertisements/:id", handler.GetAdvertisement)
}
