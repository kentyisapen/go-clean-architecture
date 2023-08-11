package file

import (
	"context"

	"github.com/kentyisapen/go-clean-architecture/models"
)

type FileRepository interface {
	CreateFile(ctx context.Context, user *models.User, fm *models.File) error
	GetFile(ctx context.Context, user *models.User) (*models.File, error)
	GetFiles(ctx context.Context, user *models.User) ([]*models.File, error)
	DeleteFile(ctx context.Context, user *models.User, id string) error
}