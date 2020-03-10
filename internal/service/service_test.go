package service

import (
	"net/http"
	"testing"

	"github.com/justindfuller/financial"
)

func TestService(t *testing.T) {
	tests := []*test{
		{
			name:       "/POST user",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostUserRequest{
				Data: &financial.User{
					Email: "service_test@example.com",
				},
			},
			response: &financial.UserResponse{},
			expected: &financial.UserResponse{
				Id: 1,
			},
		},
		{
			name:       "POST /user again",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostUserRequest{
				Data: &financial.User{
					Email: "service_test2@example.com",
				},
			},
			response: &financial.UserResponse{},
			expected: &financial.UserResponse{
				Id: 2,
			},
		},
		{
			name:       "POST /user already exists",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostUserRequest{
				Data: &financial.User{
					Email: "service_test2@example.com",
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "POST /user missing data",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostUserRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageMissingEmail,
			},
		},
		{
			name:       "POST /user missing email",
			endpoint:   endpointUser,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostUserRequest{
				Data: &financial.User{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageMissingEmail,
			},
		},
		{
			name:       "GET /user",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetUserRequest{
				Data: &financial.GetUserData{
					Email: "service_test@example.com",
				},
			},
			response: &financial.User{},
			expected: &financial.User{
				Id:    1,
				Email: "service_test@example.com",
			},
		},
		{
			name:       "GET /user not found",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetUserRequest{
				Data: &financial.GetUserData{
					Email: "not even a real email",
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "GET /user missing id",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request: &financial.GetUserRequest{
				Data: &financial.GetUserData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /user missing data",
			endpoint:   endpointUser,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request:    &financial.GetUserRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /account savings",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					Name:    "Savings",
					UserId:  1,
					Balance: 27585.45,
					Mode:    financial.Mode_INVESTMENTS,
				},
			},
			response: &financial.PostAccountResponse{},
			expected: &financial.PostAccountResponse{
				Id: 1,
			},
		},
		{
			name:       "POST /account credit card",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					Name:    "Credit Card",
					UserId:  1,
					Balance: 3496.45,
					Mode:    financial.Mode_DEBT,
				},
			},
			response: &financial.PostAccountResponse{},
			expected: &financial.PostAccountResponse{
				Id: 2,
			},
		},
		{
			name:       "POST /account already exists",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					Name:    "Credit Card",
					UserId:  1,
					Balance: 3496.45,
					Mode:    financial.Mode_DEBT,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "POST /account bad request",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /account missing name",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostAccountRequest{
				Data: &financial.Account{
					UserId: 1,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /account missing data",
			endpoint:   endpointAccount,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostAccountRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /accounts",
			endpoint:   endpointAccounts,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetAccountsRequest{
				Data: &financial.GetAccountsData{
					UserId: 1,
				},
			},
			response: &financial.GetAccountsResponse{},
			expected: &financial.GetAccountsResponse{
				Accounts: []*financial.Account{
					{
						Id:      1,
						Name:    "Savings",
						UserId:  1,
						Balance: 27585.45,
						Mode:    financial.Mode_INVESTMENTS,
					},
					{
						Id:      2,
						Name:    "Credit Card",
						UserId:  1,
						Balance: 3496.45,
						Mode:    financial.Mode_DEBT,
					},
				},
			},
		},
		{
			name:       "GET /accounts not found",
			endpoint:   endpointAccounts,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetAccountsRequest{
				Data: &financial.GetAccountsData{
					UserId: 222,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "POST /contribution credit card",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostContributionRequest{
				Data: &financial.Contribution{
					AccountId: 2,
					Amount:    500,
				},
			},
			response: &financial.PostContributionResponse{},
			expected: &financial.PostContributionResponse{
				Id: 1,
			},
		},
		{
			name:       "POST /contribution savings",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostContributionRequest{
				Data: &financial.Contribution{
					AccountId: 1,
					Amount:    500,
				},
			},
			response: &financial.PostContributionResponse{},
			expected: &financial.PostContributionResponse{
				Id: 2,
			},
		},
		{
			name:       "POST /contribution already exists",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostContributionRequest{
				Data: &financial.Contribution{
					AccountId: 1,
					Amount:    500,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "POST /contribution bad request",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostContributionRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /contribution bad request",
			endpoint:   endpointContribution,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostContributionRequest{
				Data: &financial.Contribution{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /contribution",
			endpoint:   endpointContribution,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetContributionRequest{
				Data: &financial.GetContributionData{
					AccountId: 2,
				},
			},
			response: &financial.Contribution{},
			expected: &financial.Contribution{
				Id:        1,
				AccountId: 2,
				Amount:    500,
			},
		},
		{
			name:       "GET /contribution not found",
			endpoint:   endpointContribution,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetContributionRequest{
				Data: &financial.GetContributionData{
					AccountId: 2222,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "GET /contribution missing account id",
			endpoint:   endpointContribution,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request: &financial.GetContributionRequest{
				Data: &financial.GetContributionData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /contribution missing data",
			endpoint:   endpointContribution,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request:    &financial.GetContributionRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /goal debt free",
			endpoint:   endpointGoal,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostGoalRequest{
				Data: &financial.Goal{
					Name:       "Debt Free",
					UserId:     1,
					AccountIds: []int64{1},
					Balance:    0,
				},
			},
			response: &financial.PostGoalResponse{},
			expected: &financial.PostGoalResponse{
				Id: 1,
			},
		},
		{
			name:       "POST /goal down payment",
			endpoint:   endpointGoal,
			httpMethod: http.MethodPost,
			statusCode: http.StatusCreated,
			request: &financial.PostGoalRequest{
				Data: &financial.Goal{
					Name:       "House down payment",
					UserId:     1,
					AccountIds: []int64{2},
					Balance:    75000,
				},
			},
			response: &financial.PostGoalResponse{},
			expected: &financial.PostGoalResponse{
				Id: 2,
			},
		},
		{
			name:       "POST /goal missing UserId",
			endpoint:   endpointGoal,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostGoalRequest{
				Data: &financial.Goal{
					Name:       "House down payment",
					AccountIds: []int64{2},
					Balance:    75000,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /goal missing data",
			endpoint:   endpointGoal,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request:    &financial.PostGoalRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "POST /goal duplicate name",
			endpoint:   endpointGoal,
			httpMethod: http.MethodPost,
			statusCode: http.StatusBadRequest,
			request: &financial.PostGoalRequest{
				Data: &financial.Goal{
					Name:       "House down payment",
					UserId:     1,
					AccountIds: []int64{2},
					Balance:    75000,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageAlreadyExists,
			},
		},
		{
			name:       "GET /goals",
			endpoint:   endpointGoals,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetGoalsRequest{
				Data: &financial.GetGoalData{
					UserId: 1,
				},
			},
			response: &financial.GetGoalsResponse{},
			expected: &financial.GetGoalsResponse{
				Goals: []*financial.Goal{
					{
						Name:       "Debt Free",
						UserId:     1,
						AccountIds: []int64{1},
						Balance:    0,
					},
					{
						Name:       "House down payment",
						UserId:     1,
						AccountIds: []int64{2},
						Balance:    75000,
					},
				},
			},
		},
		{
			name:       "GET /goals not found",
			endpoint:   endpointGoals,
			httpMethod: http.MethodGet,
			statusCode: http.StatusNotFound,
			request: &financial.GetGoalsRequest{
				Data: &financial.GetGoalData{
					UserId: 2,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageNotFound,
			},
		},
		{
			name:       "GET /goals missing UserId",
			endpoint:   endpointGoals,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request: &financial.GetGoalsRequest{
				Data: &financial.GetGoalData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /goals missing data",
			endpoint:   endpointGoals,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request:    &financial.GetGoalsRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /calculate",
			endpoint:   endpointCalculate,
			httpMethod: http.MethodGet,
			statusCode: http.StatusOK,
			request: &financial.GetCalculateRequest{
				Data: &financial.GetCalculateData{
					UserId:  1,
					Periods: 2,
				},
			},
			response: &financial.GetCalculateResponse{},
			expected: &financial.GetCalculateResponse{
				Periods: []*financial.Period{
					{
						Accounts: []*financial.Account{
							{
								Balance: 28085.45,
								Id:      1,
								Name:    "Savings",
								UserId:  1,
							},
							{
								Balance: 2996.45,
								Id:      2,
								Mode:    financial.Mode_DEBT,
								Name:    "Credit Card",
								UserId:  1,
							},
						},
					},
					{
						Accounts: []*financial.Account{
							{
								Balance: 28585.45,
								Id:      1,
								Name:    "Savings",
								UserId:  1,
							},
							{
								Balance: 2496.45,
								Id:      2,
								Mode:    financial.Mode_DEBT,
								Name:    "Credit Card",
								UserId:  1,
							},
						},
					},
				},
			},
		},
		{
			name:       "GET /calculate missing data",
			endpoint:   endpointCalculate,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request:    &financial.GetCalculateRequest{},
			response:   &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /calculate missing UserId",
			endpoint:   endpointCalculate,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request: &financial.GetCalculateRequest{
				Data: &financial.GetCalculateData{},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
		{
			name:       "GET /calculate missing Periods",
			endpoint:   endpointCalculate,
			httpMethod: http.MethodGet,
			statusCode: http.StatusBadRequest,
			request: &financial.GetCalculateRequest{
				Data: &financial.GetCalculateData{
					UserId: 1,
				},
			},
			response: &financial.Error{},
			expected: &financial.Error{
				Message: messageInvalidEntity,
			},
		},
	}

	runTests(t, tests)
}
