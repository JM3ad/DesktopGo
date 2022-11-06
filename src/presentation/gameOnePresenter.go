package presentation

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
	"github.com/desktop-go/src/game"
)

type GameOnePresenter struct {
	window   fyne.Window
	game     game.GameOne
	goToMenu func()
	done     chan struct{}
}

func (g *GameOnePresenter) Present() {
	quitButton := widget.NewButton("Exit", g.ReturnToMenu)
	toolbar := container.New(layout.NewHBoxLayout(), quitButton)

	header := canvas.NewText(g.game.GetName(), color.Black)
	headerBar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), header, layout.NewSpacer())

	money := widget.NewLabel(strconv.Itoa(g.game.GetMoney()))
	display := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), money, layout.NewSpacer())

	g.window.SetContent(container.New(layout.NewVBoxLayout(), toolbar, layout.NewSpacer(), headerBar, layout.NewSpacer(), display))
	quit := g.updateScreen(money)
	g.game.RunAsync(quit)
}

func (g *GameOnePresenter) ReturnToMenu() {
	fmt.Println("Exit 1 hit")
	g.goToMenu()
}

func (g *GameOnePresenter) updateScreen(money *widget.Label) chan struct{} {
	ticker := time.NewTicker(time.Second / 2)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				money.SetText(strconv.Itoa(g.game.GetMoney()))
			case <-quit:
				ticker.Stop()
				fmt.Println("Stopping the presentation")
				g.done <- struct{}{}
				return
			}
		}
	}()
	return quit
}

func (g *GameOnePresenter) End() {
	fmt.Println("Ending presenter")
	g.game.End()
}