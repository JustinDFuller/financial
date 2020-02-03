package financialcalc

import (
	"testing"
)

func TestInvestmentAccount(t *testing.T) {
	var balance Balance
	var mode ModeInvestment

	investmentAccount := &InvestmentAccount{
		BaseAccount{
			Name:                     "Investments",
			Mode:                     mode,
			Balance:                  balance.FromFloat(30000),
			InterestRate:             .055,
			AddInterestEveryNPeriods: 26,
		},
	}

	investmentContrubition := &Contribution{
		Account: investmentAccount,
		Amount:  500,
	}

	res := Calculate(&CalculateRequest{
		Contributions:  Contributions{investmentContrubition},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if res == nil || len(res) != 26 || !res[25].Accounts.Find(investmentAccount).GetBalance().Equal(45425) {
		t.Fatal("Invalid result.", res[25].Accounts.Find(investmentAccount).GetBalance().ToFloat())
	}
}

func TestDebtAccount(t *testing.T) {
	var balance Balance
	var mode ModeDebt

	debtAccount := &DebtAccount{
		BaseAccount{
			Name:                     "Auto Loan",
			Mode:                     mode,
			Balance:                  balance.FromFloat(4400),
			InterestRate:             .0264,
			AddInterestEveryNPeriods: 2,
		},
	}

	debtContribution := &Contribution{
		Account: debtAccount,
		Amount:  200,
	}

	res := Calculate(&CalculateRequest{
		Contributions:  Contributions{debtContribution},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if res == nil || len(res) != 26 || !res[21].Accounts.Find(debtAccount).GetBalance().Equal(0) {
		t.Fatal("Invalid result.", res[21].Accounts.Find(debtAccount).GetBalance().ToFloat())
	}
}
