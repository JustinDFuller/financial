package calculator

import (
	"github.com/shopspring/decimal"
)

var (
	zero   = decimal.NewFromInt(0)
	one    = decimal.NewFromInt(1)
	twelve = decimal.NewFromInt(12)
)

type (
	interestCalculator func(a *Account, periodsInvested, periodsPerYear int64) decimal.Decimal
	signGetter         func(contribution decimal.Decimal) decimal.Decimal
)

func AsInvestmentAccount(a *Account) *Account {
	a.calculateInterest = CompoundInterest
	a.getSign = positive
	a.Type = "Investment"
	return a
}

func AsDebtAccount(a *Account) *Account {
	a.calculateInterest = SimpleInterest
	a.getSign = negate
	a.Type = "Debt"
	return a
}

func (a *Account) Contribute(contribution decimal.Decimal, period int64) {
	result := a.Balance.Add(a.getSign(contribution)).RoundBank(2)

	if result.LessThanOrEqual(zero) {
		result = zero
	}

	a.Balance = result
}

func (a *Account) IsInterestPeriod(currentPeriod int64) bool {
	return currentPeriod%a.AddInterestEveryNPeriods == 0
}

func (a *Account) Is(possibleAccount *Account) bool {
	return a.Name == possibleAccount.Name
}

func (a *Account) AddInterest(currentPeriod, periodsPerYear int64) {
	a.Balance = a.calculateInterest(a, currentPeriod, periodsPerYear).RoundBank(2)
}

func (a *Account) MakeCopy() *Account {
	var newAccount Account
	newAccount = *a
	return &newAccount
}

func SimpleInterest(a *Account, periodsInvested, periodsPerYear int64) decimal.Decimal {
	return a.Balance.Mul(a.InterestRate.Add(decimal.NewFromInt(1)))
}

// A = B (1 + I / N) ^ NY
// A = New Amount
// B = Account Balance
// R = Interest Rate
// N = Number of times interest is compounded per year
// Y = Years Invested (Ignored here because we add interest once per year.)
func CompoundInterest(a *Account, periodsInvested, periodsPerYear int64) decimal.Decimal {
	b := a.Balance
	r := a.InterestRate

	x := r.Div(twelve)
	x = x.Add(one)
	x = x.Pow(twelve)
	return b.Mul(x)
}
