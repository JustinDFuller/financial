package calculator

import (
	"testing"

	"github.com/shopspring/decimal"
)

var (
	investmentAccount = AsInvestmentAccount(&Account{
		Name:                     "Investments",
		Balance:                  decimal.NewFromInt(30000),
		InterestRate:             decimal.NewFromFloat(.055),
		AddInterestEveryNPeriods: 12,
	})

	investmentContrubition = &Contribution{
		Account: investmentAccount,
		Amount:  decimal.NewFromInt(500),
	}

	debtAccount = AsDebtAccount(&Account{
		Name:                     "Auto Loan",
		Balance:                  decimal.NewFromInt(4000),
		InterestRate:             decimal.NewFromFloat(.0244),
		AddInterestEveryNPeriods: 2,
	})

	debtContribution = &Contribution{
		Account: debtAccount,
		Amount:  decimal.NewFromInt(200),
	}

	calculateRequest = &CalculateRequest{
		Contributions:  Contributions{debtContribution, investmentContrubition},
		Periods:        24,
		PeriodsPerYear: 24,
	}
)

func TestInvestmentAccount(t *testing.T) {
	periods := Calculate(calculateRequest)

	if periods == nil || len(periods) != 24 {
		t.Fatal("Not enough periods in result.", periods)
	}

	firstPeriodBalance := decimal.NewFromFloat(30500)
	if !periods.AccountBalanceAt(investmentAccount, 1).Equal(firstPeriodBalance) {
		t.Fatalf("Incorrect first period balance. Got %v: Expected: %v", periods.AccountBalanceAt(investmentAccount, 1), firstPeriodBalance)
	}

	finalBalance := decimal.NewFromFloat(46514.36)
	if !periods.AccountBalanceAt(investmentAccount, 24).Equal(finalBalance) {
		t.Fatalf("Incorrect ending balance. Got: %v Expected: %v", periods.AccountBalanceAt(investmentAccount, 24), finalBalance)
	}

	periodTwelve := periods.AccountBalanceAt(investmentAccount, 12)
	periodThirteen := periods.AccountBalanceAt(investmentAccount, 13)

	if periodThirteen.LessThan(periodTwelve) {
		t.Fatal("Balance should not go down after interest is applied.", periodTwelve, periodThirteen)
	}
}

func TestDebtAccount(t *testing.T) {
	periods := Calculate(calculateRequest)

	if periods == nil || len(periods) != 24 {
		t.Fatal("Not enough periods in result.", periods)
	}

	firstPeriodBalance := decimal.NewFromFloat(3800)
	if !periods.AccountBalanceAt(debtAccount, 1).Equal(firstPeriodBalance) {
		t.Fatalf("Incorrect first period balance. Got %v Expected %v", periods.AccountBalanceAt(debtAccount, 1), firstPeriodBalance)
	}

	finalPeriodBalance := decimal.NewFromFloat(115.1)
	if !periods.AccountBalanceAt(debtAccount, 22).Equal(finalPeriodBalance) {
		t.Fatalf("Incorrect ending balance. Got %v Expected %v", periods.At(22).Accounts.Find(debtAccount).Balance, finalPeriodBalance)
	}

	if !periods.AccountBalanceAt(debtAccount, 24).Equal(zero) {
		t.Fatal("Debt account cannot go below 0.", periods.AccountBalanceAt(debtAccount, 24))
	}
}
