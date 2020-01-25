package financialcalc

import "testing"

var (
	balance Balance

	investmentAccount = &Account{
		Name:         "Investments",
		Mode:         ModeInvestment,
		Balance:      balance.FromFloat(30000),
		InterestRate: .055,
	}

	investmentContrubition = &Contribution{
		Account: investmentAccount,
		Amount:  500,
	}
)

func TestFinancialCalculator(t *testing.T) {
	res := Calculate(&CalculateRequest{
		Contributions:  Contributions{investmentContrubition},
		Periods:        26,
		PeriodsPerYear: 26,
	})

	if res == nil || len(res) != 26 || res[25].Accounts.Find(investmentAccount).Balance.Equal(45425) {
		t.Fatal("Invalid result.", res)
	}
}
