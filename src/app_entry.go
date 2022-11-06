package src

import (
	"fmt"
	"strconv"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/desktop-go/src/game"
)

func StartApp() {
	var chans []chan struct{}
	done := make(chan struct{})
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(500, 500))

	myWidget := widget.NewLabel("Hello World!")

	w.SetContent(myWidget)

	game := game.NewGame()
	quitGame := game.RunAsync(done)
	quitDisplay := present(myWidget, game, done)

	chans = append(chans, quitGame, quitDisplay)
	defer cleanUp(chans, done)

	button := widget.NewButton("Click me", func() { fmt.Println("Tapped") })
	w.SetContent(button)

	fmt.Println("Running")
	w.ShowAndRun()
}

func present(myWidget *widget.Label, game game.Game, done chan struct{}) chan struct{} {
	ticker := time.NewTicker(time.Second / 2)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				myWidget.SetText(strconv.Itoa(game.GetMoney()))
			case <-quit:
				ticker.Stop()
				fmt.Println("Stopping the presentation")
				done <- struct{}{}
				return
			}
		}
	}()
	return quit
}

/*
Currently this doesn't really achieve anything, but was interesting to explore
*/
func cleanUp(chans []chan struct{}, done chan struct{}) {
	fmt.Println("Quitting")
	for _, channel := range chans {
		fmt.Println("Sending")
		channel <- struct{}{}
	}
	fmt.Println("Sent")
	doneCount := 0
	for {
		select {
		case <-done:
			fmt.Println("Closed channel")
			doneCount++
			if doneCount == len(chans) {
				fmt.Println("Clean up complete, exiting")
				return
			}
		case <-time.After(3 * time.Second):
			fmt.Println("Closing non-gracefully")
			return
		}
	}
}
