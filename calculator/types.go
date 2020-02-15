package calculator

import "github.com/shopspring/decimal"

type (
	User struct {
		Id       int64
		Accounts Accounts
	}

	Accounts []*Account

	Goal struct {
		Accounts []*Account
		Balance  decimal.Decimal
	}

	Contribution struct {
		Account *Account
		Amount  decimal.Decimal
	}

	Contributions []*Contribution

	CalculateRequest struct {
		Contributions  Contributions
		Periods        int64
		PeriodsPerYear int64
		CurrentPeriod  int64
	}

	Period struct {
		Accounts Accounts
	}

	Periods []Period
)
