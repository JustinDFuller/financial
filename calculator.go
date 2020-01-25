package financialcalc

func Calculate(req *CalculateRequest) []Period {
	var periods Periods

	for i := int64(0); i < req.Periods; i++ {
		periodNumber := i + 1
		period := Period{}

		for _, contribution := range req.Contributions {
			var account Account
			account.Copy(contribution.Account)
			account.Balance.Add(contribution.Amount * float64(periodNumber))

			if (periodNumber)%req.PeriodsPerYear == 0 {
				account.Balance.Compound(account.InterestRate, 12, float64(periodNumber/req.PeriodsPerYear))
			}

			period.Accounts = append(period.Accounts, account)
		}

		periods = append(periods, period)
	}

	return periods
}
