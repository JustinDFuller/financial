package financialcalc

import "math"

type Account interface {
	AddInterest(periodsInvested, periodsPerYear float64)
	Contribute(contribution, period float64)
	MakeCopy() Account
	IsInterestPeriod(currentPeriod float64) bool
	Is(possibleAccount Account) bool
	GetName() string
	GetBalance() *Balance
}

type (
	BaseAccount struct {
		Name                     string
		Balance                  *Balance
		Mode                     Mode
		InterestRate             float64
		AddInterestEveryNPeriods float64
	}

	InvestmentAccount struct {
		BaseAccount
	}

	DebtAccount struct {
		BaseAccount
	}
)

func (account *BaseAccount) Contribute(contribution, period float64) {
	account.Balance.Add(account.Mode.GetContribution(contribution))
}

func (account *BaseAccount) IsInterestPeriod(currentPeriod float64) bool {
	return math.Mod(currentPeriod, account.AddInterestEveryNPeriods) == 0
}

func (account *BaseAccount) Is(possibleAccount Account) bool {
	return account.GetName() == possibleAccount.GetName()
}

func (account *BaseAccount) GetName() string {
	return account.Name
}

func (account *BaseAccount) GetBalance() *Balance {
	return account.Balance
}

func (account *InvestmentAccount) AddInterest(periodsInvested, periodsPerYear float64) {
	const numberTimesAddInterested = 12
	exponent := math.Pow(1+(account.InterestRate/numberTimesAddInterested), numberTimesAddInterested*(periodsPerYear/periodsInvested))
	account.Balance.FromFloat(math.Floor(account.Balance.ToFloat() * exponent))
}

func (account *InvestmentAccount) MakeCopy() Account {
	var accountCopy InvestmentAccount
	accountCopy = *account
	*accountCopy.Balance = *account.Balance
	return &accountCopy
}

func (account *DebtAccount) AddInterest(periodsInvested, periodsPerYear float64) {
	account.Balance.FromFloat(math.Floor(account.Balance.ToFloat() * (1 + account.InterestRate)))
}

func (account *DebtAccount) MakeCopy() Account {
	var accountCopy DebtAccount
	accountCopy = *account
	*accountCopy.Balance = *account.Balance
	return &accountCopy
}
