package financialcalc

func (account *Account) CopyFrom(account2 *Account) {
	*account = *account2
}

func (account *Account) Contribute(contribution float64) {
	account.Balance.Add(account.Mode.GetContribution(contribution))
}
