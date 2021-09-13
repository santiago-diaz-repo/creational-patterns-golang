package template

import "fmt"

type ISocialMedia interface {
	Login(username string, password string) bool
	Logout()
}

type socialMedia struct {
	socialMedia ISocialMedia
	username    string
	password    string
}

func (sm socialMedia) Post(message string) {
	if sm.socialMedia.Login(sm.username, sm.password) {
		fmt.Println("posting..." + message)
		sm.socialMedia.Logout()
	} else {
		fmt.Println("Bye")
	}
}
