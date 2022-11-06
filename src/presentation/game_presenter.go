package presentation

type GamePresenter interface {
	Present()
	ReturnToMenu()
	End()
}
