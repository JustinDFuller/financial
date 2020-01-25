package financialcalc

func (account *Account) Copy(account2 *Account) {
	*account = *account2
}
