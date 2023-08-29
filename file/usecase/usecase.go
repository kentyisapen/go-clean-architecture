package usecase

import (
	"context"
	"os"

	"github.com/kentyisapen/go-clean-architecture/file"
	"github.com/kentyisapen/go-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileUseCase struct {
	fileRepo file.Repository
}

func NewFileUseCase(fileRepo file.Repository) *FileUseCase {
	return &FileUseCase{
		fileRepo: fileRepo,
	}
}

func (f FileUseCase) CreateFile(ctx context.Context, user *models.User, name string, folderId string, bin []byte) error {
	id := primitive.NewObjectID().Hex()
	// TODO: ファイル文字数とか容量とかでバリデーションする
	// TODO: ファイル保存先をMinIOとかにしたい
	fp, err := os.Create("storage/" + id)
	if err != nil {
		return file.ErrFaildToCreateFile
	}

	_, err = fp.Write(bin)
	if err != nil {
		return file.ErrFaildToCreateFile
	}

	fm := &models.File{
		ID:       id,
		Name:     name,
		FolderID: folderId,
		UserID:   user.ID,
	}

	return f.fileRepo.CreateFile(ctx, user, fm)
}
