package decorator

import "fmt"

type basicWebPage struct {
}

func NewBasicWebPage() WebSite {
	return basicWebPage{}
}

func (b basicWebPage) Display() {
	fmt.Println("Basic site")
}

func (b basicWebPage) Close() string {
	return "Exiting.."
}
