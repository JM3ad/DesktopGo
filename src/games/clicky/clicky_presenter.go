package clicky

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

type ClickyPresenter struct {
	Window      fyne.Window
	Game        Clicky
	GoToMenu    func()
	Done        chan struct{}
	moneyWidget *widget.Label
}

func (g *ClickyPresenter) Present() {
	homeTab := g.getHomeTab()
	storeTab := g.getStoreTab()

	// Two tabs
	tabs := container.NewAppTabs(
		homeTab,
		storeTab,
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	g.Window.SetContent(tabs)

	quit := g.updateScreen()
	g.Game.RunAsync(quit)
}

func (g *ClickyPresenter) ReturnToMenu() {
	fmt.Println("Exit 1 hit")
	g.GoToMenu()
}

func (g *ClickyPresenter) updateScreen() chan struct{} {
	ticker := time.NewTicker(time.Second / 10)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				g.moneyWidget.SetText(strconv.Itoa(g.Game.GetMoney()))
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

func (g *ClickyPresenter) End() {
	fmt.Println("Ending presenter")
	g.Game.End()
}

func (g *ClickyPresenter) getHomeTab() *container.TabItem {

	quitButton := widget.NewButton("Exit", g.ReturnToMenu)
	toolbar := container.New(layout.NewHBoxLayout(), quitButton)

	header := canvas.NewText(g.Game.GetName(), color.Black)
	headerBar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), header, layout.NewSpacer())

	money := widget.NewLabel(strconv.Itoa(g.Game.GetMoney()))
	g.moneyWidget = money
	clickForMoneyButton := widget.NewButton("Click Me!", g.Game.Click)
	display := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), money, clickForMoneyButton, layout.NewSpacer())

	fullScreen := container.New(layout.NewVBoxLayout(), toolbar, layout.NewSpacer(), headerBar, layout.NewSpacer(), display)
	return container.NewTabItem("Home", fullScreen)
}

func (g *ClickyPresenter) getStoreTab() *container.TabItem {
	listOfUpgrades := GetUpgrades()
	upgradeElements := make([]fyne.CanvasObject, len(listOfUpgrades))
	for i, upgrade := range listOfUpgrades {
		// force closure
		upgrade := upgrade
		cost := widget.NewLabel(strconv.Itoa(upgrade.cost))
		buy := widget.NewButton("Buy", func() { g.Game.BuyUpgrade(upgrade) })
		display := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), cost, buy, layout.NewSpacer())

		upgradeElements[i] = display
	}

	fullScreen := container.New(layout.NewVBoxLayout(), upgradeElements...)

	return container.NewTabItem("Store", fullScreen)
}
