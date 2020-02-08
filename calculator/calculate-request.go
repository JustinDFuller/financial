package service

func (r *CalculateRequest) NextPeriod() bool {
	r.CurrentPeriod += 1

	return r.CurrentPeriod <= r.Periods
}
