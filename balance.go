package financialcalc

import "math"

func (balance *Balance) ToFloat() float64 {
	return float64(*balance)
}

func (balance *Balance) FromFloat(f float64) *Balance {
	b := Balance(f)
	balance = &b
	return balance
}

func (balance *Balance) Add(contribution float64) {
	balance.FromFloat(balance.ToFloat() + contribution)
}

func (balance *Balance) Compound(interestRate, numberTimesCompounded, timeInvested float64) {
	exponent := math.Pow(1+(interestRate/numberTimesCompounded), numberTimesCompounded*timeInvested)
	balance.FromFloat(math.Floor(balance.ToFloat() * exponent))
}

func (balance *Balance) Equal(f float64) bool {
	return balance.ToFloat() == f
}
