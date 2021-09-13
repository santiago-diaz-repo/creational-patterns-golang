package template

import "fmt"

type Facebook struct {
	sm socialMedia
}

func (f Facebook) Login(username string, password string) bool {
	f.sm.username = username
	f.sm.password = password
	fmt.Println("login from facebook")
	return true
}

func (f Facebook) Logout() {
	fmt.Println("logout from facebook")
}
