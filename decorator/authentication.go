package decorator

import "fmt"

type Authentication struct {
	WebPageDecorator
}

func NewAuthentication(website WebSite) WebSite {
	return Authentication{
		WebPageDecorator{
			wrappee: website,
		},
	}
}

func (a Authentication) Display() {
	a.WebPageDecorator.wrappee.Display()
	a.Authenticate()
}

func (a Authentication) Close() string {
	panic("implement me")
}

func (a Authentication) Authenticate() {
	fmt.Println("Authenticating..")
}
