package template

import "fmt"

type Twitter struct {
	sm socialMedia
}

func (t Twitter) Login(username string, password string) bool {
	t.sm.username = username
	t.sm.password = password
	fmt.Println("login from facebook")
	return true
}

func (t Twitter) Logout() {
	fmt.Println("logout from facebook")
}
