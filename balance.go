package financialcalc

import (
	"fmt"
	"math"
)

func (balance *Balance) ToFloat() float64 {
	return float64(*balance)
}

func (balance *Balance) FromFloat(f float64) *Balance {
	*balance = Balance(f)
	return balance
}

func (balance *Balance) Add(contribution float64) {
	amount := balance.ToFloat() + contribution

	if amount < 0 {
		amount = 0
	}

	balance.FromFloat(amount)
}

// numberTimesCompounded * periodsInvest is WRONG
// numberTimesCompounded * (periodsPerYear / periodsInvest) is CORRECT
func (balance *Balance) Compound(interestRate, periodsInvested, periodsPerYear float64) {
	const numberTimesCompounded = 12
	exponent := math.Pow(1+(interestRate/numberTimesCompounded), numberTimesCompounded*(periodsPerYear/periodsInvested))
	balance.FromFloat(math.Floor(balance.ToFloat() * exponent))
}

func (balance *Balance) Equal(f float64) bool {
	return balance.ToFloat() == f
}

// String implements the fmt.Stringer interface.
// This makes it easy to drop into fmt.Print(account.Balance)
func (balance *Balance) String() string {
	return fmt.Sprintf("%v", balance.ToFloat())
}
