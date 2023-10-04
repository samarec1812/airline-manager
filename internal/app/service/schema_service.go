package service

import (
	"context"

	"github.com/samarec1812/airline-manager/internal/app/entity"
)

func (a *app) CreateSchema(ctx context.Context, schema entity.Schema) error {
	return a.schemaRepo.Create(ctx, schema)
}

func (a *app) RemoveSchema(ctx context.Context, schemaID int64) error {
	return a.schemaRepo.Remove(ctx, schemaID)
}
