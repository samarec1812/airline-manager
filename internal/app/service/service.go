package service

import (
	"context"

	entity "github.com/samarec1812/airline-manager/internal/app/entity/airline"
)

type AirlineRepository interface {
	Create(context.Context, entity.Airline) error
	Remove(context.Context, string) error
}

type App interface {
	CreateAirline(ctx context.Context, airline entity.Airline) error
	RemoveAirline(ctx context.Context, code string) error
}

type app struct {
	airlineRepo AirlineRepository
}

func NewApp(airlineRepo AirlineRepository) App {
	return &app{
		airlineRepo: airlineRepo,
	}
}

func (a *app) CreateAirline(ctx context.Context, airline entity.Airline) error {
	return a.airlineRepo.Create(ctx, airline)
}

func (a *app) RemoveAirline(ctx context.Context, code string) error {
	return a.airlineRepo.Remove(ctx, code)
}
