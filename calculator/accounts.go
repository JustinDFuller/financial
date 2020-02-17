package calculator

func (accounts Accounts) Find(account *Account) *Account {
	for _, possibleAccount := range accounts {
		if account.Is(possibleAccount) {
			return possibleAccount
		}
	}

	return nil
}

func (accounts Accounts) Append(account *Account) Accounts {
	return append(accounts, account)
}
