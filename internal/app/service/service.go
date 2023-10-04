package service

import (
	"context"
	"github.com/samarec1812/airline-manager/internal/app/entity"
)

type AirlineRepository interface {
	Create(context.Context, entity.Airline) error
	Remove(context.Context, string) error
}

type AccountRepository interface {
	Create(context.Context, entity.Account) error
	Remove(context.Context, int64) error
}

type SchemaRepository interface {
	Create(context.Context, entity.Schema) error
	Remove(context.Context, int64) error
}

type App interface {
	// Airline methods
	CreateAirline(ctx context.Context, airline entity.Airline) error
	RemoveAirline(ctx context.Context, code string) error

	// Account methods
	CreateAccount(ctx context.Context, account entity.Account) error
	RemoveAccount(ctx context.Context, accountID int64) error

	// Schema methods
	CreateSchema(ctx context.Context, schema entity.Schema) error
	RemoveSchema(ctx context.Context, schemaID int64) error
}

type app struct {
	airlineRepo AirlineRepository
	accountRepo AccountRepository
	schemaRepo  SchemaRepository
}

func NewApp(airlineRepo AirlineRepository, accountRepo AccountRepository, schemaRepo SchemaRepository) App {
	return &app{
		airlineRepo: airlineRepo,
		accountRepo: accountRepo,
		schemaRepo:  schemaRepo,
	}
}
