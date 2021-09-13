package chainOfResponsability

import "testing"

func TestUserNameProvider_IsAuthorize(t *testing.T) {
	prov := OauthProvider{
		token: "kjjbjhvhj",
	}

	proc := UsernameProcessor{}
	proc2 := OauthProcessor{}
	proc.SetNext(&proc2)

	proc.IsAuthorized(&prov)
}
