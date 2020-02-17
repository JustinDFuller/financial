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

	debtGoal = &Goal{
		Name: "Debt Free",
		Accounts: Accounts{
			debtAccount,
		},
		Balance: zero,
	}

	calculateRequest = &CalculateRequest{
		Contributions:  Contributions{debtContribution, investmentContrubition},
		Goals:          Goals{debtGoal},
		Periods:        24,
		PeriodsPerYear: 24,
	}
)

func TestInvestmentAccount(t *testing.T) {
	periods := Calculate(calculateRequest)

	if periods == nil || len(periods) != 24 {
		t.Fatal("Not enough periods in result.", periods)
	}

	expected := decimal.NewFromFloat(30500)
	actual := periods.AccountBalanceAt(investmentAccount, 1)
	if !actual.Equal(expected) {
		t.Fatalf("Incorrect first period balance. Got %v: Expected: %v", actual, expected)
	}

	expected = decimal.NewFromFloat(46514.36)
	actual = periods.AccountBalanceAt(investmentAccount, 24)
	if !actual.Equal(expected) {
		t.Fatalf("Incorrect ending balance. Got: %v Expected: %v", actual, expected)
	}

	periodTwelve := periods.AccountBalanceAt(investmentAccount, 12)
	periodThirteen := periods.AccountBalanceAt(investmentAccount, 13)
	if periodThirteen.LessThan(periodTwelve) {
		t.Fatal("Balance should not go down after interest is applied.", periodTwelve, periodThirteen)
	}

	expected = decimal.NewFromFloat(3800)
	actual = periods.AccountBalanceAt(debtAccount, 1)
	if !actual.Equal(expected) {
		t.Fatalf("Incorrect first period balance. Got %v Expected %v", actual, expected)
	}

	expected = decimal.NewFromFloat(115.1)
	actual = periods.AccountBalanceAt(debtAccount, 22)
	if !actual.Equal(expected) {
		t.Fatalf("Incorrect ending balance. Got %v Expected %v", actual, expected)
	}

	actual = periods.AccountBalanceAt(debtAccount, 24)
	if !actual.Equal(zero) {
		t.Fatal("Debt account cannot go below 0.", actual)
	}

	if goal := periods.GoalAt(debtGoal, 23); goal == nil {
		t.Fatal("Expected Goal to be met in period 23. Got nil.")
	}

	if goal := periods.GoalAt(debtGoal, 24); goal != nil {
		t.Fatal("Goal should not appear more than once. Found goal after period 23.")
	}
}
