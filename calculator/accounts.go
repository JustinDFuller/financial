package calculator

import "github.com/shopspring/decimal"

func (accounts Accounts) Find(account *Account) *Account {
	for _, possibleAccount := range accounts {
		if account.Is(possibleAccount) {
			return possibleAccount
		}
	}

	return nil
}

func (accounts Accounts) AccountBalanceEqual(account *Account, amount decimal.Decimal) bool {
	return accounts.Find(account).Balance.Equal(amount)
}
