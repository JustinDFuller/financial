package financialcalc

func (account *Account) CopyFrom(account2 *Account) {
	*account = *account2
}
