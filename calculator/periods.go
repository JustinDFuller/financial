package calculator

import "github.com/shopspring/decimal"

func (p Periods) Last() *Period {
	if len(p) == 0 {
		return nil
	}
	return p[len(p)-1]
}

func (p Periods) At(period int) *Period {
	return p[period-1]
}

func (p Periods) AccountBalanceAt(account *Account, period int) decimal.Decimal {
	return p.At(period).Accounts.Find(account).Balance
}
