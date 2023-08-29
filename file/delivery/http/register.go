package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kentyisapen/go-clean-architecture/file"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc file.UseCase) {
	h := NewHandler(uc)

	bookmarks := router.Group("/files")
	{
		bookmarks.POST("", h.Create)
	}
}
