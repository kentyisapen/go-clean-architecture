package file

import (
	"bytes"
	"context"

	"github.com/kentyisapen/go-clean-architecture/models"
)

type FileUsecase interface {
	CreateFile(ctx context.Context, user *models.User, filename string, bin bytes.Buffer) error
	GetFile(ctx context.Context, user *models.User) (*models.File, error)
	GetFiles(ctx context.Context, user *models.User) ([]*models.File, error)
	DeleteFile(ctx context.Context, user *models.User, id string) error
}
