package chainOfResponsability

import "fmt"

type UserNameProvider struct {
	username string
	password string
}

func (u *UserNameProvider) Validate() bool {
	fmt.Println("Username provider")
	return false
}
