package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kentyisapen/go-clean-architecture/file"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc file.UseCase) {
	h := NewHandler(uc)

	files := router.Group("/files")
	{
		files.POST("", h.Create)
		files.GET("/:id", h.Get)
	}
}
