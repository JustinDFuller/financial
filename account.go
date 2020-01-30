package financialcalc

import "math"

type Account struct {
	Name                  string
	Balance               *Balance
	Mode                  Mode
	InterestRate          float64
	CompoundEveryNPeriods float64
}

func (account *Account) CopyFrom(account2 *Account) {
	*account = *account2
	*account.Balance = *account2.Balance
}

func (account *Account) Contribute(contribution, period float64) {
	account.Balance.Add(account.Mode.GetContribution(contribution))
}

func (account *Account) IsCompoundingPeriod(currentPeriod float64) bool {
	return math.Mod(currentPeriod, account.CompoundEveryNPeriods) == 0
}
