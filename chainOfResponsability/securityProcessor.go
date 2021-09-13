package chainOfResponsability

type SecurityProcessor interface {
	IsAuthorized(provider SecProvider) bool
	SetNext(next SecurityProcessor)
}
