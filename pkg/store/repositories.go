package store

import (
	"context"

	"github.com/Feruz666/advertising/pkg/model"
)

type PostRepository interface {
	Save(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	FindAll(ctx context.Context) ([]model.PostModel, error)
	FindOneTrue(ctx context.Context, id int, fl bool) (*model.PostModel, error)
	FindOneFalse(ctx context.Context, id int, fl bool) (*model.PostModel, error)
}