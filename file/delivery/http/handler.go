package http

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kentyisapen/go-clean-architecture/auth"
	"github.com/kentyisapen/go-clean-architecture/file"
	"github.com/kentyisapen/go-clean-architecture/models"
)

type File struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FolderID string `json:"folder_id"`
}

type Handler struct {
	useCase file.UseCase
}

func NewHandler(useCase file.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type createInput struct {
	Name     string `json:"name"`
	FolderId string `json:"folder_id"`
	Bin      string `json:"bin"`
}

func (h *Handler) Create(c *gin.Context) {
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	bin, err := base64.StdEncoding.DecodeString(inp.Bin)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateFile(c.Request.Context(), user, inp.Name, inp.FolderId, bin); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type GetResponse struct {
	ID       string `json:"id"`
	FolderID string `json:"folder_id"`
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	file, err := h.useCase.GetFile(c.Request.Context(), user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, &GetResponse{
		ID:       file.ID,
		FolderID: file.FolderID,
		UserId:   file.UserID,
		Name:     file.Name,
	})
}

type IndexResponse struct {
	Files []*models.File `json:"files"`
}

func (h *Handler) Index(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	files, err := h.useCase.GetFiles(c.Request.Context(), user)

	if err != nil {

		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, &IndexResponse{
		Files: files,
	})
}
