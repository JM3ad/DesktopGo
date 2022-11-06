package clicky

import (
	"testing"
)

func TestClickingAddsMoney(t *testing.T) {
	game := NewClicky()
	original_money := game.GetMoney()
	game.Click()
	money := game.GetMoney()
	if money != original_money+1 {
		t.Errorf("Output %q not equal to %q", money, original_money)
	}
}

func TestBuyingUpgradeDoesNothingIfNotEnoughMoney(t *testing.T) {
	game := NewClicky()
	upgrade := getTestUpgrade()
	if game.money >= upgrade.cost {
		t.Errorf("Game has enough money to start with, test doesn't make sense")
	}

	game.BuyUpgrade(&upgrade)

	if game.stats.moneyPerClick != 1 {
		t.Errorf("Money per click %d has changed", game.stats.moneyPerClick)
	}
}

func TestBuyingUpgradeImprovesStats(t *testing.T) {
	game := NewClicky()
	upgrade := getTestUpgrade()
	game.money = upgrade.cost

	game.BuyUpgrade(&upgrade)
	if game.stats.moneyPerClick != 2 {
		t.Errorf("Money per click %d not equal to 2", game.stats.moneyPerClick)
	}
	if game.money != 0 {
		t.Errorf("Money hasn't been spent")
	}
}

func TestMoneyPerClickApplied(t *testing.T) {
	game := NewClicky()
	moneyPerClick := 5
	game.stats.moneyPerClick = moneyPerClick
	game.money = 0

	game.Click()

	if game.GetMoney() != moneyPerClick {
		t.Errorf("Money: %d hasn't increased by correct amount %d", game.GetMoney(), moneyPerClick)
	}
}

func getTestUpgrade() Upgrade {
	return Upgrade{cost: 10, effect: func(stats *ClickyStats) { stats.moneyPerClick += 1 }}
}
