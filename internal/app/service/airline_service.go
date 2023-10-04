package service

import (
	"context"

	"github.com/samarec1812/airline-manager/internal/app/entity"
)

func (a *app) CreateAirline(ctx context.Context, airline entity.Airline) error {
	return a.airlineRepo.Create(ctx, airline)
}

func (a *app) RemoveAirline(ctx context.Context, code string) error {
	return a.airlineRepo.Remove(ctx, code)
}
