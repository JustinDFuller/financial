package service

import "context"

type db interface {
	PutUser(context.Context, *User) error
	GetUser(context.Context, int64) (*User, error)
	PutUserAccount(context.Context, *Account) error
	GetUserAccounts(context.Context, *User) ([]*Account, error)
}
