package calculator

func Calculate(req *CalculateRequest) Periods {
	var periods Periods

	for req.NextPeriod() {
		var period Period

		for _, contribution := range req.Contributions {
			var account *Account

			if lastPeriod := periods.Last(); lastPeriod != nil {
				account = lastPeriod.Accounts.Find(contribution.Account)
			}

			if account == nil {
				account = contribution.Account
			}

			account = account.MakeCopy()
			account.Contribute(contribution.Amount, req.CurrentPeriod)

			if account.IsInterestPeriod(req.CurrentPeriod) {
				account.AddInterest(req.CurrentPeriod, req.PeriodsPerYear)
			}

			period.Accounts = append(period.Accounts, account)
		}

		periods = append(periods, &period)
	}

	req.Reset()
	return periods
}
