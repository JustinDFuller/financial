package financialcalc

func (accounts Accounts) Find(account *Account) *Account {
	for _, possibleAccount := range accounts {
		if possibleAccount.Name == account.Name {
			return &possibleAccount
		}
	}

	return nil
}
