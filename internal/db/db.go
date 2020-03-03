package db

import (
	"errors"

	"github.com/justindfuller/financial"
)

var (
	ErrMissingEmail  = errors.New("Missing email.")
	ErrAlreadyExists = errors.New("Entity already exists.")
	ErrNotFound      = errors.New("Not found.")
)

type Store interface {
	CreateUserByEmail(email string) (int64, error)
	GetUserByEmail(email string) (*financial.UserResponse, error)
	CreateAccountByUserId(userid int64, account *financial.Account) (int64, error)
	GetAccountsByUserId(userid int64) ([]*financial.Account, error)
}
