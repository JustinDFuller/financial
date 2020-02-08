package service

import (
	"math"

	"github.com/justindfuller/financial-calculator/types"
)

type InterestCalculator func(a *account, periodsInvested, periodsPerYear float64)
type ContributeByMode func(a *account, contribution float64)

type account struct {
	types.Account
	ContibuteByMode   ContributeByMode
	CalculateInterest InterestCalculator
}

func AsInvestmentAccount(a *account) *Account {
	a.CalculateInterest = CompoundInterest
	return account
}

func AsDebtAccount(a *account) *Account {
	a.CalculateInterest = SimpleInterest
	return account
}

func (a *account) Contribute(contribution, period float64) {
	a.Balance.Add(a.Mode.GetContribution(contribution))
}

func (a *account) IsInterestPeriod(currentPeriod float64) bool {
	return math.Mod(currentPeriod, a.AddInterestEveryNPeriods) == 0
}

func (a *account) Is(possibleAccount *Account) bool {
	return a.GetName() == possibleAccount.GetName()
}

func (a *account) GetName() string {
	return a.Name
}

func (a *account) GetBalance() *Balance {
	return a.Balance
}

func (a *account) AddInterest(currentPeriod, periodsPerYear float64) {
	a.CalculateInterest(account, currentPeriod, periodsPerYear)
}

func (a *account) MakeCopy() *Account {
	var newAccount Account
	newAccount = *account
	*newAccount.Balance = *a.Balance
	return &newAccount
}

func SimpleInterest(a *account, periodsInvested, periodsPerYear float64) {
	a.Balance.FromFloat(math.Floor(a.Balance.ToFloat() * (1 + a.InterestRate)))
}

func CompoundInterest(a *account, periodsInvested, periodsPerYear float64) {
	const numberTimesAddInterested = 12
	exponent := math.Pow(1+(a.InterestRate/numberTimesAddInterested), numberTimesAddInterested*(periodsPerYear/periodsInvested))
	a.Balance.FromFloat(math.Floor(a.Balance.ToFloat() * exponent))
}
