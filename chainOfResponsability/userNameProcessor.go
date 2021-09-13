package chainOfResponsability

import "fmt"

type UsernameProcessor struct {
	next     SecurityProcessor
	username string
	password string
}

func (up *UsernameProcessor) IsAuthorized(provider SecProvider) bool {
	fmt.Println("Username processor")
	switch provider.(type) {
	case *UserNameProvider:
		provider.Validate()
	default:
		return up.next.IsAuthorized(provider)
	}
	return false
}

func (up *UsernameProcessor) SetNext(next SecurityProcessor) {
	up.next = next
	fmt.Println()
}
