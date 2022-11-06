package clicky

type ClickyStats struct {
	moneyPerClick int
	moneyPerTick  int
}

func NewClickyStats() ClickyStats {
	return ClickyStats{
		moneyPerClick: 1,
		moneyPerTick:  0,
	}
}
