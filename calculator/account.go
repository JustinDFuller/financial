package calculator

import (
	"github.com/shopspring/decimal"
)

var (
	zero   = decimal.NewFromInt(0)
	one    = decimal.NewFromInt(1)
	twelve = decimal.NewFromInt(12)
)

type InterestCalculator func(a *Account, periodsInvested, periodsPerYear int64) decimal.Decimal
type ContributeByMode func(contribution decimal.Decimal) decimal.Decimal

type Account struct {
	Name                     string
	Type                     string
	Balance                  decimal.Decimal
	InterestRate             decimal.Decimal
	AddInterestEveryNPeriods int64
	GetSign                  ContributeByMode
	calculateInterest        InterestCalculator
}

func AsInvestmentAccount(a *Account) *Account {
	a.calculateInterest = CompoundInterest
	a.GetSign = positive
	a.Type = "Investment"
	return a
}

func AsDebtAccount(a *Account) *Account {
	a.calculateInterest = SimpleInterest
	a.GetSign = negate
	a.Type = "Debt"
	return a
}

func (a *Account) Contribute(contribution decimal.Decimal, period int64) {
	result := a.Balance.Add(a.GetSign(contribution)).RoundBank(2)

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
