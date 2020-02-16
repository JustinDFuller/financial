package calculator

func (r *CalculateRequest) NextPeriod() bool {
	r.CurrentPeriod += 1

	return r.CurrentPeriod <= r.Periods
}

func (r *CalculateRequest) Reset() {
	r.CurrentPeriod = 0
}
