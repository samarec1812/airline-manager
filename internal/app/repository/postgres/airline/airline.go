package airline

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/samarec1812/airline-manager/internal/app/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewAirlineRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, airline entity.Airline) error {
	const query = `INSERT INTO airline (code, name) VALUES ($1, $2) RETURNING code`

	err := r.db.QueryRowContext(ctx, query, airline.Code, airline.Name).Scan(&airline.Code)
	if err != nil {
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
