package game_one

import (
	"fmt"
	"time"
)

type GameOne struct {
	money int
	quit  chan struct{}
}

func NewGameOne() *GameOne {
	return &GameOne{money: 0}
}

func (g *GameOne) RunAsync(done chan struct{}) {
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

func (g *GameOne) Update() {
	g.money += 1
}

func (g *GameOne) GetMoney() int {
	return g.money
}

func (g *GameOne) GetName() string {
	return "Game 1"
}

func (g *GameOne) End() {
	g.quit <- struct{}{}
}
