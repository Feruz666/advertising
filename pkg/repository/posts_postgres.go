package repository

import (
	"github.com/Feruz666/advertising/pkg/model"
	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsRepository(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (p *PostsPostgres) Save(post *model.PostModel) (*model.PostModel, error) {
	return nil, nil
}

func (p *PostsPostgres) FindAll() ([]model.PostModel, error) {
	return nil, nil
}

func (p *PostsPostgres) FindOneTrue(id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}

func (p *PostsPostgres) FindOneFalse(id int, fl bool) (*model.PostModel, error) {
	return nil, nil
}
