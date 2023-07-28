package store

import (
	"context"

	"github.com/Feruz666/advertising/pkg/model"
	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsRepository(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (p *PostsPostgres) Save(ctx context.Context, post *model.PostModel) (*model.PostModel, error) {

	return nil, nil
}

func (p *PostsPostgres) FindAll(ctx context.Context, ) ([]model.PostModel, error) {
	return nil, nil
}

func (p *PostsPostgres) FindOneTrue(ctx context.Context, id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}

func (p *PostsPostgres) FindOneFalse(ctx context.Context, id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}
