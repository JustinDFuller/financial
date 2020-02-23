package calculator

import "github.com/shopspring/decimal"

func negate(contribution decimal.Decimal) decimal.Decimal {
	return contribution.Neg()
}

func positive(contribution decimal.Decimal) decimal.Decimal {
	return contribution
}
