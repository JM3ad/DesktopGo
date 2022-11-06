package game_two

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type GameTwoPresenter struct {
	Window   fyne.Window
	Game     GameTwo
	GoToMenu func()
	Done     chan struct{}
}

func (g *GameTwoPresenter) Present() {
	quitButton := widget.NewButton("Exit", g.ReturnToMenu)
	toolbar := container.New(layout.NewHBoxLayout(), quitButton)

	header := canvas.NewText(g.Game.GetName(), color.Black)
	headerBar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), header, layout.NewSpacer())

	money := widget.NewLabel(strconv.Itoa(g.Game.GetMoney()))
	display := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), money, layout.NewSpacer())

	g.Window.SetContent(container.New(layout.NewVBoxLayout(), toolbar, layout.NewSpacer(), headerBar, layout.NewSpacer(), display))
	quit := g.updateScreen(money)
	g.Game.RunAsync(quit)
}

func (g *GameTwoPresenter) ReturnToMenu() {
	g.GoToMenu()
}

func (g *GameTwoPresenter) updateScreen(money *widget.Label) chan struct{} {
	ticker := time.NewTicker(time.Second / 2)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				money.SetText(strconv.Itoa(g.Game.GetMoney()))
			case <-quit:
				ticker.Stop()
				fmt.Println("Stopping the presentation")
				g.Done <- struct{}{}
				return
			}
		}
	}()
	return quit
}

func (g *GameTwoPresenter) End() {
	g.Game.End()
}
