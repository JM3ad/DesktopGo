package game

type Game interface {
	RunAsync(chan struct{})
	GetName() string
	End()
}
