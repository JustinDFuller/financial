package financialcalc

type Account struct {
	Name         string
	Balance      *Balance
	Mode         Mode
	InterestRate float64
}

func (account *Account) CopyFrom(account2 *Account) {
	var balance Balance
	account.Balance = balance.FromFloat(account2.Balance.ToFloat())
	account.Name = account2.Name
	account.Mode = account2.Mode
	account.InterestRate = account2.InterestRate
}

func (account *Account) Contribute(contribution, period float64) {
	account.Balance.Add(account.Mode.GetContribution(contribution * period))
}
