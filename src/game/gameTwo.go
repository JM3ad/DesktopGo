package game

import (
	"fmt"
	"time"
)

type GameTwo struct {
	money int
	quit  chan struct{}
}

func NewGameTwo() *GameTwo {
	return &GameTwo{money: 0}
}

func (g *GameTwo) RunAsync(done chan struct{}) {
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

func (g *GameTwo) Update() {
	g.money -= 1
}

func (g *GameTwo) GetMoney() int {
	return g.money
}

func (g *GameTwo) GetName() string {
	return "Game 2"
}

func (g *GameTwo) End() {
	g.quit <- struct{}{}
}
