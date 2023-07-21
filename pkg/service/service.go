package service

import (
	"github.com/Feruz666/advertising/pkg/model"
	"github.com/Feruz666/advertising/pkg/repository"
)

type PostService interface {
	Validate(post *model.PostModel) error
	Create(post *model.PostModel) (*model.PostModel, error)
	FindAll() (*model.PostModel, error)
	FindOneTrue(id int, fl bool) (*model.PostModel, error)
}

type Service struct {
	PostService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		PostService: NewPostService(repos.PostsRepository),
	}
}
