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
	// User
	CreateUserByEmail(email string) (int64, error)
	GetUserByEmail(email string) (*financial.User, error)

	// Account
	CreateAccountByUserId(userId int64, data *financial.Account) (int64, error)
	GetAccountsByUserId(userId int64) ([]*financial.Account, error)

	// Contribution
	CreateContributionByAccountId(accountId int64, data *financial.Contribution) (int64, error)
	GetContributionByAccountId(accountId int64) (*financial.Contribution, error)

	// Goal
	CreateGoalByUserId(userId int64, data *financial.Goal) (int64, error)
	GetGoalsByUserId(userId int64) ([]*financial.Goal, error)
}
