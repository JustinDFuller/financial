# Financial Calculator

[![Build Status](https://cloud.drone.io/api/badges/JustinDFuller/financial/status.svg)](https://cloud.drone.io/JustinDFuller/financial)
[![Go Report Card](https://goreportcard.com/badge/github.com/JustinDFuller/financial)](https://goreportcard.com/report/github.com/JustinDFuller/financial)

## The Idea

1. I want to make a financial calculator app.
2. I should be able to give it all my accounts + their current balance.
3. I should be able to says how much I save to those accounts each period
    Note: A period could be a week, every 2 weeks, the first and 15th, monthly, whatever.
      This allows the calculator to work for people with different paycheck types.
      I think individual accounts should also be able to get different periods.
      For example, a roth-ira might get $6k per year, rather than a monthly contribution.
4. I should also be able to track my debt (mortgage, car payment, etc.)
    Note: Debts would have a current balance, interest rate, and monthly payment.
5. I should be able to set goals. A goal should be able to have one or more accounts associated with it.
    Note: A goal is met when the combination of accounts reaches the goal level.
          A goal should be able to have a debt account associated with it, which would subtract from the
          other accounts. So, for example, a pay off mortgage might have a savings account + mortgage account.
          The goal might be $0, meaning that the savings account can pay off the mortgage.
6. I should be able to play out the calculations over an arbitrary number of periods.
7. I should be able to change contributions/payments at any time.
    For example, if I pay off my car, I should be able to up my savings contribution at that point.
8. Do I also need to specify how often to compound interest?
    No: It would be better to specify how many periods per year.

### Scenario 1:

* I am 27 years old. I want to retire when I am 50.
* I get paid bi-weekly, every 2 weeks, or 26 times a year.
* 23 years * 26 pay-periods per year = 598 pay periods.
* Current accounts:
  * Roth IRA:
      Balance, $34,000; Contribution: $230; Interest: 5.5%.
  * Investments:
      Balance, $30,200; Contribution: $500; Interest: 5.5%.
  * 529 (1):
      Balance, $9,500; Contribution: $100; Interest: 5.5%.
  * 529 (2):
      Balance, $6,000; Contribution: $100: Interest: 5.5%.
  * 401k:
      Balance, $38,300; Contribution: $650; Interest: 5.5%.
  * Acorns:
      Balance, $1000; Contribution: $100; Interest: 0%.
  * Emergency Savings:
      Balance, $14,200; Contribution: 200; Interest: 1.8%.
  * Mortgage:
      Balance, $148,750; Contribution: $525; Interest: 4.375%.
  * Auto Loan:
      Balance, $4,400; Contribution: $200; Interest: 2.64%.
* Goals:
  * Retirement:
      X amount in Roth IRA and 401k
  * College, kid 1
      X amount in 529 (1)
  * College, kid 2
      X amount in 529 (2)
  * Buy a house
      X amount in Investments
  * Emergency fund
      X amount in Emergency Savings
  * Debt Free
      Mortgage and Auto Loan to $0 balance
  * Some kind of house project
      X amount in Acorns
