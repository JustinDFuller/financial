package calculator

import "github.com/shopspring/decimal"

func (p *Period) CalculateGoals(goals Goals) {
	for _, goal := range goals {
		if goal.completed {
			continue
		}

		var balance decimal.Decimal

		for _, account := range goal.Accounts {
			if found := p.Accounts.Find(account); found != nil {
				balance = balance.Add(found.getSign(found.Balance))
			}
		}

		if goal.IsMet(balance) {
			goal.completed = true
			p.Goals = append(p.Goals, goal)
		}
	}
}
