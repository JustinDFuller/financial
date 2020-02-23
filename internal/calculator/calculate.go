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
			account.Contribute(contribution.Amount)

			if account.IsInterestPeriod(req.CurrentPeriod) {
				account.AddInterest(req.CurrentPeriod, req.PeriodsPerYear)
			}

			period.Accounts = period.Accounts.Append(account)
		}

		period.CalculateGoals(req.Goals)
		periods = periods.Append(&period)
	}

	return periods
}
