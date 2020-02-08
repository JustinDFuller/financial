package service

type Mode interface {
	GetContribution(float64) float64
}

type ModeDebt struct{}

type ModeInvestment struct{}

func (m ModeDebt) GetContribution(contribution float64) float64 {
	return -contribution
}

func (m ModeInvestment) GetContribution(contribution float64) float64 {
	return contribution
}
