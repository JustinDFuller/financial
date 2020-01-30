package financialcalc

func Calculate(req *CalculateRequest) Periods {
	var periods Periods

	for req.NextPeriod() {
		var period Period

		for _, contribution := range req.Contributions {
			var account Account
			account.CopyFrom(contribution.Account)
			account.Contribute(contribution.Amount, req.CurrentPeriod)

			if account.IsCompoundingPeriod(req.CurrentPeriod) {
				account.Balance.Compound(account.InterestRate, req.CurrentPeriod, req.PeriodsPerYear)
			}

			period.Accounts = append(period.Accounts, account)
		}

		periods = append(periods, period)
	}

	return periods
}
