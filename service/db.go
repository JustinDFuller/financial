package service

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type db interface {
	PutUser(context.Context, *User) error
	GetUser(context.Context, int64) (*User, error)
	PutUserAccount(context.Context, *Account) error
	GetUserAccounts(context.Context, *User) (Accounts, error)
}

type ds struct {
	client *datastore.Client
}

func NewDatastoreDb(ctx context.Context, cfg *Config) (db, error) {
	client, err := datastore.NewClient(ctx, cfg.GoogleCloudProject)
	if err != nil {
		return nil, fmt.Errorf("Error initializing datastore client. %w", err)
	}

	return &ds{client: client}, nil
}

func (db *ds) PutUser(ctx context.Context, user *User) error {
	return nil
}

func (db *ds) GetUser(ctx context.Context, id int64) (*User, error) {
	return nil, nil
}

func (db *ds) PutUserAccount(ctx context.Context, account *Account) error {
	return nil
}

func (db *ds) GetUserAccounts(ctx context.Context, user *User) (Accounts, error) {
	return nil, nil
}
