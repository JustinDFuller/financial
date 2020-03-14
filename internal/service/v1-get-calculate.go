package service

import (
	context "context"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server/kit"
	"github.com/gogo/protobuf/proto"
	"github.com/justindfuller/financial"
	"github.com/justindfuller/financial/internal/calculator"
	"github.com/shopspring/decimal"
)

func decodeUserCalculate(ctx context.Context, req *http.Request) (interface{}, error) {
	var request financial.GetCalculateRequest

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	err = proto.Unmarshal(body, &request)
	if err != nil {
		return nil, kit.NewProtoStatusResponse(&financial.Error{Message: err.Error()}, http.StatusBadRequest)
	}

	return &request, nil
}

// getUserCalculate has really complicated logic due to the fact that I have to convert between the
// financial and calculator types.
// I need to either
// 	1. Use the same types or
//  2. Create a cleaner way to convert back and forth between the types.
func (s *service) getUserCalculate(ctx context.Context, request interface{}) (response interface{}, err error) {
	r := request.(*financial.GetCalculateRequest)

	if r.Data == nil || r.Data.UserId == 0 || r.Data.Periods == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil
	}

	var fContributions []*financial.Contribution
	fAccounts, _ := s.db.GetAccountsByUserId(r.Data.UserId)
	fGoals, _ := s.db.GetGoalsByUserId(r.Data.UserId)

	if len(fAccounts) == 0 {
		return kit.NewProtoStatusResponse(&financial.Error{
			Message: messageInvalidEntity,
		}, http.StatusBadRequest), nil
	}

	for _, account := range fAccounts {
		contribution, _ := s.db.GetContributionByAccountId(account.Id)
		fContributions = append(fContributions, contribution)
	}

	var cGoals calculator.Goals
	var cContributions calculator.Contributions
	cAccountsById := map[int64]*calculator.Account{}
	accountsByName := map[string]*financial.Account{}
	goalsByName := map[string]*financial.Goal{}

	for _, account := range fAccounts {
		cAccount := &calculator.Account{
			Name:                     account.Name,
			Type:                     account.Mode.String(),
			Balance:                  decimal.NewFromFloat(account.Balance),
			InterestRate:             decimal.NewFromFloat(account.InterestRate),
			AddInterestEveryNPeriods: account.AddInterestEveryNPeriods,
		}

		switch account.Mode {
		case financial.Mode_INVESTMENTS:
			cAccount = calculator.AsInvestmentAccount(cAccount)
		case financial.Mode_DEBT:
			cAccount = calculator.AsDebtAccount(cAccount)
		}

		cAccountsById[account.Id] = cAccount
		accountsByName[account.Name] = account
	}

	for _, contribution := range fContributions {
		cContributions = append(cContributions, &calculator.Contribution{
			Amount:  decimal.NewFromFloat(contribution.Amount),
			Account: cAccountsById[contribution.AccountId],
		})
	}

	for _, goal := range fGoals {
		cGoal := &calculator.Goal{
			Name:     goal.Name,
			Accounts: calculator.Accounts{},
			Balance:  decimal.NewFromFloat(goal.Balance),
		}

		for _, accountId := range goal.AccountIds {
			cGoal.Accounts = append(cGoal.Accounts, cAccountsById[accountId])
		}

		cGoals = append(cGoals, cGoal)
		goalsByName[goal.Name] = goal
	}

	cReq := &calculator.CalculateRequest{
		Goals:          cGoals,
		Contributions:  cContributions,
		Periods:        r.Data.Periods,
		PeriodsPerYear: 26,
	}
	periods := calculator.Calculate(cReq)

	res := &financial.GetCalculateResponse{
		Periods: []*financial.Period{},
	}

	for _, period := range periods {
		fPeriod := &financial.Period{
			Accounts: []*financial.Account{},
		}

		for _, account := range period.Accounts {
			balance, _ := account.Balance.Float64()

			var fAccount financial.Account
			fAccount = *accountsByName[account.Name]
			fAccount.Balance = balance
			fPeriod.Accounts = append(fPeriod.Accounts, &fAccount)
		}

		for _, goal := range period.Goals {
			balance, _ := goal.Balance.Float64()

			var fGoal financial.Goal
			fGoal = *goalsByName[goal.Name]
			fGoal.Balance = balance
			fPeriod.Goals = append(fPeriod.Goals, &fGoal)
		}

		res.Periods = append(res.Periods, fPeriod)
	}

	return kit.NewProtoStatusResponse(res, http.StatusOK), nil
}
