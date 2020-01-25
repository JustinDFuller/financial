package financialcalc

func (c *Contribution) ForPeriod(period int64) float64 {
	return c.Amount * float64(period)
}
