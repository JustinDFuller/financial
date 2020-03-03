package memory

import (
	"github.com/justindfuller/financial"
	"github.com/justindfuller/financial/internal/db"
)

func New() db.Store {
	return &memory{
		accountsByUserId: map[int64][]*financial.Account{},
		usersByEmail:     map[string]*financial.UserResponse{},
	}
}

type memory struct {
	accountId        int64
	accountsByUserId map[int64][]*financial.Account

	userId       int64
	usersByEmail map[string]*financial.UserResponse
}

func (s *memory) CreateUserByEmail(email string) (int64, error) {
	if email == "" {
		return 0, db.ErrMissingEmail
	}

	if _, ok := s.usersByEmail[email]; ok {
		return 0, db.ErrAlreadyExists
	}

	s.userId++
	s.usersByEmail[email] = &financial.UserResponse{
		Id:    s.userId,
		Email: email,
	}
	return s.userId, nil
}

func (s *memory) GetUserByEmail(email string) (*financial.UserResponse, error) {
	if user, ok := s.usersByEmail[email]; ok {
		return user, nil
	}

	return nil, db.ErrNotFound
}

func (s *memory) CreateAccountByUserId(userId int64, data *financial.Account) (int64, error) {
	s.accountId++
	data.Id = s.accountId

	if accounts, ok := s.accountsByUserId[data.UserId]; !ok {
		s.accountsByUserId[data.UserId] = []*financial.Account{data}
	} else {
		s.accountsByUserId[data.UserId] = append(accounts, data)
	}

	return data.Id, nil
}

func (s *memory) GetAccountsByUserId(userId int64) ([]*financial.Account, error) {
	if accounts, ok := s.accountsByUserId[userId]; ok {
		return accounts, nil
	}

	return nil, db.ErrNotFound
}
