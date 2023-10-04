package account

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/samarec1812/airline-manager/internal/app/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, account entity.Account) error {
	const query = `INSERT INTO account (schema_id) VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, account.SchemaID).Scan(&account.ID)
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) Remove(ctx context.Context, accountID int64) error {
	const query = `DELETE FROM account WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, accountID)
	if err != nil {
		return err
	}

	return nil

}
