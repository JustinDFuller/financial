package calculator

import "github.com/shopspring/decimal"

func (g *Goal) IsMet(balance decimal.Decimal) bool {
	return balance.GreaterThanOrEqual(g.Balance)
}
