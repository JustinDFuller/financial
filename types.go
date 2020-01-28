package financialcalc

type (
	Balance float64

	Account struct {
		Name         string
		Balance      *Balance
		Mode         Mode
		InterestRate float64
	}

	Accounts []Account

	Goal struct {
		Accounts []*Account
		Balance  float64
	}

	Contribution struct {
		Account *Account
		Amount  float64
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
