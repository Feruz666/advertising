package repository

import (
	"github.com/Feruz666/advertising/pkg/model"
	"github.com/jmoiron/sqlx"
)

type PostsRepository interface {
	Save(post *model.PostModel) (*model.PostModel, error)
	FindAll() ([]model.PostModel, error)
	FindOneTrue(id int, fl bool) (*model.PostModel, error)
	FindOneFalse(id int, fl bool) (*model.PostModel, error)
}

type Repository struct {
	PostsRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PostsRepository: NewPostsRepository(db),
	}
}
