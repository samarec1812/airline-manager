package postgres

import (
	"fmt"

	// Register postgres driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(connection string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
