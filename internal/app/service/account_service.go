package service

import (
	"context"

	"github.com/samarec1812/airline-manager/internal/app/entity"
)

func (a *app) CreateAccount(ctx context.Context, account entity.Account) error {
	return a.accountRepo.Create(ctx, account)
}

func (a *app) RemoveAccount(ctx context.Context, accountID int64) error {
	return a.accountRepo.Remove(ctx, accountID)
}
