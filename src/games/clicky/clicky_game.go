package clicky

import (
	"fmt"
	"time"
)

type Clicky struct {
	money int
	quit  chan struct{}
	stats ClickyStats
}

func NewClicky() *Clicky {
	return &Clicky{money: 0, stats: NewClickyStats()}
}

func (g *Clicky) RunAsync(done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	g.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Updating")
				g.Update()
			case <-g.quit:
				ticker.Stop()
				fmt.Println("Stopping the game loop")
				done <- struct{}{}
				return
			}
		}
	}()
}

func (g *Clicky) Update() {
	g.money += g.stats.moneyPerTick
}

func (g *Clicky) Click() {
	g.money += g.stats.moneyPerClick
}

func (g *Clicky) GetMoney() int {
	return g.money
}

func (g *Clicky) GetName() string {
	return "Clicky"
}

func (g *Clicky) BuyUpgrade(upgrade *Upgrade) {
	fmt.Println("Trying to buy")
	fmt.Printf("%d vs %d", g.money, upgrade.cost)
	if g.money >= upgrade.cost {
		g.money -= upgrade.cost
		upgrade.effect(&g.stats)
	}
}

func (g *Clicky) End() {
	g.quit <- struct{}{}
}
