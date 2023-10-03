package airline

import (
	"context"
	"database/sql"

	entity "github.com/samarec1812/airline-manager/internal/app/entity/airline"
)

const (
	airlineTable = "users"
)

type Repository struct {
	db *sql.DB
}

func NewAirlineRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, airline entity.Airline) error {
	const query = `INSERT INTO airline (code, name) VALUES ($1, $2) RETURNING code`
	if err := r.db.QueryRowContext(ctx, query, airline.Code, airline.Name).Scan(&airline.Code); err != nil {
		return err
	}

	return nil

}

func (r *Repository) Remove(ctx context.Context, code string) error {
	const query = `DELETE FROM airline WHERE code = $1`
	_, err := r.db.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	return nil

}
