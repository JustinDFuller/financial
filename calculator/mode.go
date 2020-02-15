package calculator

import "github.com/shopspring/decimal"

func debtContributor(contribution decimal.Decimal) decimal.Decimal {
	return contribution.Neg()
}

func investmentContributor(contribution decimal.Decimal) decimal.Decimal {
	return contribution
}
