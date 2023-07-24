package service

import (
	"context"
	"errors"

	"github.com/Feruz666/advertising/pkg/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	PostService
}

func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		PostService: NewPostService(ctx, store),
	}, nil
}