package service

import (
	"context"

	"github.com/Feruz666/advertising/pkg/model"
	"github.com/Feruz666/advertising/pkg/store"
)

type PostsService struct {
	ctx context.Context
	store *store.Store
}

func NewPostService(ctx context.Context, store *store.Store) *PostsService {
	return &PostsService{ctx: ctx, store: store}
}

func (p *PostsService) Validate(ctx context.Context, post *model.PostModel) error {
	return nil
}

func (p *PostsService) Create(ctx context.Context, post *model.PostModel) (*model.PostModel, error) {
	
	return nil, nil
}

func (p *PostsService) FindAll(ctx context.Context, ) (*model.PostModel, error) {
	return nil, nil
}

func (p *PostsService) FindOneTrue(ctx context.Context, id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}
