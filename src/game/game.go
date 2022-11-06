package game

import (
	"fmt"
	"time"
)

type Game interface {
	GetMoney() int
	RunAsync(chan struct{}) chan struct{}
}

type game struct {
	money int
}

func NewGame() Game {
	return &game{money: 0}
}

func (g *game) RunAsync(done chan struct{}) chan struct{} {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Updating")
				g.Update()
			case <-quit:
				ticker.Stop()
				fmt.Println("Stopping the game loop")
				done <- struct{}{}
				return
			}
		}
	}()
	return quit
}

func (g *game) Update() {
	g.money += 1
}

func (g *game) GetMoney() int {
	return g.money
}
