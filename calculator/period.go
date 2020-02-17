package calculator

import "github.com/shopspring/decimal"

func (p *Period) CalculateGoals(goals Goals) {
	for _, goal := range goals {
		var balance decimal.Decimal

		for _, account := range goal.Accounts {
			if found := p.Accounts.Find(account); found != nil {
				balance = balance.Add(found.GetSign(found.Balance))
			}
		}

		if balance.Equal(goal.Balance) {
			p.Goals = append(p.Goals, goal)
		}
	}
}
