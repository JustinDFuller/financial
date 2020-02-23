package calculator

func (g Goals) Find(goal *Goal) *Goal {
	for _, possibleGoal := range g {
		if possibleGoal.Name == goal.Name {
			return possibleGoal
		}
	}

	return nil
}
