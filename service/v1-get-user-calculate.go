package service

import (
	context "context"
	"net/http"

	"github.com/justindfuller/financial/calculator"
	"github.com/shopspring/decimal"
)

func decodeUserCalculate(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func (s *service) getUserCalculate(ctx context.Context, request interface{}) (response interface{}, err error) {
	periods := calculator.Calculate(&calculator.CalculateRequest{
		Contributions: calculator.Contributions{
			&calculator.Contribution{
				Account: calculator.AsInvestmentAccount(&calculator.Account{
					Name:                     "Investments",
					Balance:                  decimal.NewFromInt(30000),
					InterestRate:             decimal.NewFromFloat(0.55),
					AddInterestEveryNPeriods: 26,
				}),
				Amount: decimal.NewFromInt(500),
			},
		},
		Periods:        52,
		PeriodsPerYear: 52,
	})

	return periods, nil
}
