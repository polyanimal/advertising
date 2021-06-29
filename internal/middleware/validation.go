package middleware

import "github.com/gin-gonic/gin"

type Validation interface {
	CheckValid() gin.HandlerFunc
}

type ValidationMiddleware struct {
}

func (m *ValidationMiddleware) CheckValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
