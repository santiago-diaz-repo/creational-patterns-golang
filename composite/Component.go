package composite

type Component interface {
	Play()
	Stop()
	getName() string
}
