package chainOfResponsability

import "fmt"

type OauthProvider struct {
	token string
}

func (o *OauthProvider) Validate() bool {
	fmt.Println("Oauth provider")
	return true
}
