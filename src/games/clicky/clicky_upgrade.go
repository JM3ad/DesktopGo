package clicky

type Upgrade struct {
	cost   int
	effect func(*ClickyStats)
}

func GetUpgrades() []*Upgrade {
	return []*Upgrade{
		{cost: 10, effect: oneMorePerClick},
		{cost: 25, effect: oneMorePerTick},
		{cost: 50, effect: doublePerTick},
	}
}

func oneMorePerClick(stats *ClickyStats) {
	stats.moneyPerClick += 1
}

func oneMorePerTick(stats *ClickyStats) {
	stats.moneyPerTick += 1
}

func doublePerTick(stats *ClickyStats) {
	stats.moneyPerTick *= 2
}
