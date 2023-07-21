package service

import (
	"github.com/Feruz666/advertising/pkg/model"
	"github.com/Feruz666/advertising/pkg/repository"
)

type PostsService struct {
	repo repository.PostsRepository
}

func NewPostService(repo repository.PostsRepository) *PostsService {
	return &PostsService{repo: repo}
}

func (p *PostsService) Validate(post *model.PostModel) error {
	return nil
}

func (p *PostsService) Create(post *model.PostModel) (*model.PostModel, error) {
	
	return nil, nil
}

func (p *PostsService) FindAll() (*model.PostModel, error) {
	return nil, nil
}

func (p *PostsService) FindOneTrue(id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}
