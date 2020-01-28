package financialcalc

func Calculate(req *CalculateRequest) []Period {
	var periods Periods

	for req.NextPeriod() {
		var period Period

		for _, contribution := range req.Contributions {
			var account Account
			account.CopyFrom(contribution.Account)
			account.Contribute(contribution.Amount)

			if req.IsCompoundingPeriod() {
				account.Balance.Compound(account.InterestRate, req.PeriodsInvested())
			}

			period.Accounts = append(period.Accounts, account)
		}

		periods = append(periods, period)
	}

	return periods
}
