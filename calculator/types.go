package calculator

import "github.com/shopspring/decimal"

type (
	User struct {
		Id       int64
		Accounts Accounts
	}

	Accounts []*Account

	Goal struct {
		Name     string
		Accounts Accounts
		Balance  decimal.Decimal
	}

	Goals []*Goal

	Contribution struct {
		Account *Account
		Amount  decimal.Decimal
	}

	Contributions []*Contribution

	CalculateRequest struct {
		Contributions  Contributions
		Goals          Goals
		Periods        int64
		PeriodsPerYear int64
		CurrentPeriod  int64
	}

	Period struct {
		Accounts Accounts
		Goals    Goals
	}

	Periods []*Period
)
