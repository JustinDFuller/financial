package financialcalc

type Mode interface {
	GetContribution(float64) float64
}

type ModeDebt string

type ModeInvestment string

func (m ModeDebt) GetContribution(contribution float64) float64 {
	return -contribution
}

func (m ModeInvestment) GetContribution(contribution float64) float64 {
	return contribution
}
