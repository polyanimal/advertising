package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/advertising/internal/models"
	constants "github.com/polyanimal/advertising/pkg/const"
	"github.com/polyanimal/advertising/pkg/util"
	"net/http"
)

type Validation interface {
	CheckValid() gin.HandlerFunc
}

type ValidationMiddleware struct {
}

func NewValidationMiddleware() *ValidationMiddleware{
	return &ValidationMiddleware{}
}

func (m *ValidationMiddleware) CheckValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invalid := false
		errMsg := ""
		ad := new(models.Advertisement)
		err := ctx.BindJSON(ad)
		if err != nil {
			util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if ad.Name == "" || ad.Description == ""  || len(ad.PhotoLinks) == 0 {
			err := ctx.Error(errors.New("invalid advertisement fields"))
			if err != nil {
				util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
				return
			}
			ctx.Status(http.StatusBadRequest)
			return
		}

		if len(ad.Name) > 200 {
			invalid = true
			errMsg += "advertisement name is too long (max 200) "

		}

		if len(ad.Description) > 1000 {
			invalid = true
			errMsg += "advertisement description is too long (max 1000) "
		}

		if len(ad.PhotoLinks) > 3 {
			invalid = true
			errMsg += "too many photos (max 3)"

		}

		if invalid {
			util.RespondWithError(ctx, http.StatusBadRequest, errMsg)
			return
		}

		ctx.Set(constants.NewAdKey, ad)
		ctx.Next()
	}
}
