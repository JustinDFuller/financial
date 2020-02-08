package service

func Calculate(req *CalculateRequest) Periods {
	var periods Periods

	for req.NextPeriod() {
		var period Period

		for _, contribution := range req.Contributions {
			account := contribution.Account.MakeCopy()
			account.Contribute(contribution.Amount, req.CurrentPeriod)

			if account.IsInterestPeriod(req.CurrentPeriod) {
				account.AddInterest(req.CurrentPeriod, req.PeriodsPerYear)
			}

			period.Accounts = append(period.Accounts, account)
		}

		periods = append(periods, period)
	}

	return periods
}
