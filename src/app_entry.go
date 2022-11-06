package src

import (
	"fmt"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/desktop-go/src/presentation"
)

func StartApp() {
	var chans []chan struct{}
	done := make(chan struct{})
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(1000, 600))

	presentation := presentation.NewPresentation(w, done)

	chans = append(chans, presentation.Quit)
	defer cleanUp(chans, done)

	presentation.Start()

	fmt.Println("Running")
	w.ShowAndRun()
}

/*
Currently this doesn't really achieve anything, but was interesting to explore
*/
func cleanUp(chans []chan struct{}, done chan struct{}) {
	fmt.Println("Quitting")
	for _, channel := range chans {
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
