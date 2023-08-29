package mongo

import (
	"context"
	"fmt"

	"github.com/kentyisapen/go-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   primitive.ObjectID `bson:"userId"`
	Name     string             `bson:"name"`
	FolderId primitive.ObjectID `bson:"folderId,omitempty"`
}

type FileRepository struct {
	db *mongo.Collection
}

func NewFileRepository(db *mongo.Database, collection string) *FileRepository {
	return &FileRepository{
		db: db.Collection(collection),
	}
}

func (r FileRepository) CreateFile(ctx context.Context, user *models.User, fm *models.File) error {
	fm.UserID = user.ID

	model := toModel(fm)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		fmt.Println("konohendesukane", err)
		return err
	}

	fm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func toModel(f *models.File) *File {
	uid, _ := primitive.ObjectIDFromHex(f.UserID)
	fid, _ := primitive.ObjectIDFromHex(f.FolderID) // Errが起きたらNilObjectID

	return &File{
		UserId:   uid,
		FolderId: fid,
		Name:     f.Name,
	}
}