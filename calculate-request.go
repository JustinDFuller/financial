package financialcalc

func (r *CalculateRequest) IsCompoundingPeriod() bool {
	return (r.CurrentPeriod)%r.PeriodsPerYear == 0
}

func (r *CalculateRequest) NextPeriod() bool {
	r.CurrentPeriod += 1

	return r.CurrentPeriod <= r.Periods
}

func (r *CalculateRequest) PeriodsInvested() int64 {
	return r.CurrentPeriod / r.PeriodsPerYear
}
