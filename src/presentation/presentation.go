package presentation

import (
	"fmt"
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/desktop-go/src/games/clicky"
	"github.com/desktop-go/src/games/game_two"
)

type Presentation struct {
	window           fyne.Window
	Quit             chan struct{}
	done             chan struct{}
	currentPresenter GamePresenter
}

func NewPresentation(w fyne.Window, done chan struct{}) *Presentation {
	return &Presentation{
		w,
		make(chan struct{}, 1),
		done,
		nil,
	}
}

func (p *Presentation) Start() {
	p.ShowMenu()
	go p.CleanUp()
}

func (p *Presentation) ShowMenu() {
	fmt.Println("Showing")
	ClickyButton := widget.NewButton("Play Game 1", func() { p.startG1() })
	gameTwoButton := widget.NewButton("Play Game 2", func() { p.startG2() })
	content := container.New(layout.NewVBoxLayout(), ClickyButton, gameTwoButton)

	header := canvas.NewText("Welcome!", color.Black)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), header, layout.NewSpacer())

	p.window.SetContent(container.New(layout.NewVBoxLayout(), centered, layout.NewSpacer(), content, layout.NewSpacer()))
}

func (p *Presentation) CleanUp() {
	for {
		select {
		case <-p.Quit:
			fmt.Println("Cleaning")
			p.currentPresenter.End()
			p.done <- struct{}{}
			return
		}
	}
}

func (p *Presentation) startG1() {
	Clicky := clicky.NewClicky()
	p.currentPresenter = &clicky.ClickyPresenter{
		Window:   p.window,
		Game:     *Clicky,
		GoToMenu: p.goToMenu,
		Done:     p.done,
	}
	p.currentPresenter.Present()
}

func (p *Presentation) startG2() {
	gameTwo := game_two.NewGameTwo()
	p.currentPresenter = &game_two.GameTwoPresenter{
		Window:   p.window,
		Game:     *gameTwo,
		GoToMenu: p.goToMenu,
		Done:     p.done,
	}
	p.currentPresenter.Present()
}

func (p *Presentation) goToMenu() {
	fmt.Println("goToMenu hit")
	if p.currentPresenter != nil {
		p.currentPresenter.End()
	}
	fmt.Println("Here")
	p.ShowMenu()
}
