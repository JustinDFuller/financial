package calculator

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestInvestmentAccount(t *testing.T) {

	investmentAccount := AsInvestmentAccount(&Account{
		Name:                     "Investments",
		Balance:                  decimal.NewFromInt(30000),
		InterestRate:             decimal.NewFromFloat(.055),
		AddInterestEveryNPeriods: 26,
	})

	investmentContrubition := &Contribution{
		Account: investmentAccount,
		Amount:  decimal.NewFromInt(500),
	}

	periods := Calculate(&CalculateRequest{
		Contributions:  Contributions{investmentContrubition},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if periods == nil || len(periods) != 26 || !periods[25].Accounts.AccountBalanceEqual(investmentAccount, decimal.NewFromFloat(45425.54)) {
		t.Fatal("Invalid result.", periods[25].Accounts.Find(investmentAccount).Balance)
	}
}

func TestDebtAccount(t *testing.T) {
	debtAccount := AsDebtAccount(&Account{
		Name:                     "Auto Loan",
		Balance:                  decimal.NewFromInt(4400),
		InterestRate:             decimal.NewFromFloat(.0264),
		AddInterestEveryNPeriods: 2,
	})

	debtContribution := &Contribution{
		Account: debtAccount,
		Amount:  decimal.NewFromInt(200),
	}

	periods := Calculate(&CalculateRequest{
		Contributions:  Contributions{debtContribution},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if periods == nil || len(periods) != 26 || !periods[21].Accounts.Find(debtAccount).Balance.Equal(decimal.NewFromInt(0)) {
		t.Fatal("Invalid result.", periods[21].Accounts.Find(debtAccount).Balance)
	}
}
