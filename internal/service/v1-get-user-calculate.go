package service

import (
	context "context"
	"net/http"

	"github.com/justindfuller/financial/internal/calculator"
	"github.com/shopspring/decimal"
)

func decodeUserCalculate(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func (s *service) getUserCalculate(ctx context.Context, request interface{}) (response interface{}, err error) {
	investmentAccount := calculator.AsInvestmentAccount(&calculator.Account{
		Name:                     "Investments",
		Balance:                  decimal.NewFromInt(30000),
		InterestRate:             decimal.NewFromFloat(0.055),
		AddInterestEveryNPeriods: 26,
	})

	debtAccount := calculator.AsDebtAccount(&calculator.Account{
		Name:                     "Car Loan",
		Balance:                  decimal.NewFromInt(4000),
		InterestRate:             decimal.NewFromFloat(0.0375),
		AddInterestEveryNPeriods: 2,
	})

	periods := calculator.Calculate(&calculator.CalculateRequest{
		Goals: calculator.Goals{
			&calculator.Goal{
				Name: "Debt Free",
				Accounts: calculator.Accounts{
					debtAccount,
				},
				Balance: decimal.Decimal{},
			},
			&calculator.Goal{
				Name: "House down payment",
				Accounts: calculator.Accounts{
					investmentAccount,
				},
				Balance: decimal.NewFromInt(74000),
			},
		},
		Contributions: calculator.Contributions{
			&calculator.Contribution{
				Account: investmentAccount,
				Amount:  decimal.NewFromInt(500),
			},
			&calculator.Contribution{
				Account: debtAccount,
				Amount:  decimal.NewFromInt(400),
			},
		},
		Periods:        598,
		PeriodsPerYear: 26,
	})

	return periods, nil
}
