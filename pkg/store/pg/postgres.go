package pg

import (
	"time"

	"github.com/Feruz666/advertising/configs"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a postgres DB
type DB struct {
	*sqlx.DB
}


// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg := configs.Get()

	pgDB, err := sqlx.Connect(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		return nil, err
	}

	err = pgDB.Ping()
	if err != nil {
		return nil, err
	}

	pgDB.SetConnMaxIdleTime(time.Second * Timeout)

	return &DB{pgDB}, nil
}	