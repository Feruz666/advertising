package service

import (
	"context"

	"github.com/Feruz666/advertising/pkg/model"
)

type PostService interface {
	Validate(ctx context.Context, post *model.PostModel) error
	Create(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	FindAll(ctx context.Context, ) (*model.PostModel, error)
	FindOneTrue(ctx context.Context, id int, fl bool) (*model.PostModel, error)
}



// func NewService(repos *store.Repository) *Service {
// 	return &Service{
// 		PostService: NewPostService(repos.PostRepository),
// 	}
// }
