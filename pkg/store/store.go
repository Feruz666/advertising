package store

import (
	"context"
	"fmt"

	"github.com/Feruz666/advertising/pkg/store/pg"
	"github.com/jmoiron/sqlx"
)

// Store contains all repositories
type Store struct {
	Pg *sqlx.DB
	PostRepository
}

// New creates new store
func New(ctx context.Context) (*Store, error) {
	// cfg := configs.Get()

	
	// connect to Postgres
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, fmt.Errorf("pgDB.Dial failed, error - %w", err)
	}


	var store Store

	// Init Postgres repositories
	if pgDB != nil {
		store.Pg = pgDB.DB
	}


	return &store, nil

}