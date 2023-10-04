package schema

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/samarec1812/airline-manager/internal/app/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewSchemaRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, schema entity.Schema) error {
	const query = `INSERT INTO schema (name) VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, schema.Name).Scan(&schema.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Remove(ctx context.Context, schemaID int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	var tmp int64
	const selectQuery = `SELECT 1 FROM account WHERE schema_id = $1 LIMIT 1`

	err = tx.GetContext(ctx, &tmp, selectQuery, schemaID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to select")
	}
	if tmp > 0 {
		tx.Rollback()
		return fmt.Errorf("exist account with this schema_id")
	}

	const query = `DELETE FROM schema WHERE id = $1`

	_, err = tx.ExecContext(ctx, query, schemaID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
