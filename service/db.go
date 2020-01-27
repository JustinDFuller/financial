package service

import "context"

type db interface {
	PutUser(context.Context, *types.User) error
	GetUser(context.Context, int64) (*types.User, error)
	PutUserAccount(context.Context, *types.Account) error
	GetUserAccounts(context.Context, *types.User) ([]*types.Account, error)
}
