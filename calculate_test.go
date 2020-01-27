package financialcalc

import (
	"testing"
)

func TestInvestmentAccount(t *testing.T) {
	var balance Balance

	investmentAccount := &Account{
		Name:         "Investments",
		Mode:         ModeInvestment,
		Balance:      balance.FromFloat(30000),
		InterestRate: .055,
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

	if res == nil || len(res) != 26 || !res[25].Accounts.Find(investmentAccount).Balance.Equal(45425) {
		t.Fatal("Invalid result.", res[25].Accounts.Find(investmentAccount).Balance.ToFloat())
	}
}

func TestDebtAccount(t *testing.T) {
	var balance Balance

	debtAccount := &Account{
		Name:         "Auto Loan",
		Mode:         ModeDebt,
		Balance:      balance.FromFloat(4400),
		InterestRate: .0264,
	}

	contribution := &Contribution{
		Account: debtAccount,
		Amount:  200,
	}

	res := Calculate(&CalculateRequest{
		Contributions:  Contributions{contribution},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if res == nil || len(res) != 26 || !res[25].Accounts.Find(debtAccount).Balance.Equal(0) {
		t.Fatal("Invalid result.", res[25].Accounts.Find(debtAccount).Balance.ToFloat())
	}
}
