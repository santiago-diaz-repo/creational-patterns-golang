package chainOfResponsability

import "fmt"

type OauthProcessor struct {
	next SecurityProcessor
}

func (op *OauthProcessor) IsAuthorized(provider SecProvider) bool {
	fmt.Println("Oauth processor")
	switch provider.(type) {
	case *OauthProvider:
		provider.Validate()
	default:
		return op.next.IsAuthorized(provider)
	}
	return false
}

func (op *OauthProcessor) SetNext(next SecurityProcessor) {
	op.next = next
}
