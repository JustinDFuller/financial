package financialcalc

import (
	"testing"
)

func TestInvestmentAccount(t *testing.T) {
	var balance Balance
	var mode ModeInvestment

	investmentAccount := &Account{
		Name:         "Investments",
		Mode:         mode,
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
	var mode ModeDebt

	debtAccount := &Account{
		Name:         "Auto Loan",
		Mode:         mode,
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

	// Something needs to keep the balance from going below 0.
	// Who does it? Should the contribution have a periods field?
	// Should balance.Add() disallow negative balances?
	// Should account use Mode to see if it can go negative?
	if res == nil || len(res) != 26 || !res[25].Accounts.Find(debtAccount).Balance.Equal(0) {
		t.Fatal("Invalid result.", res[25].Accounts.Find(debtAccount).Balance.ToFloat())
	}
}
