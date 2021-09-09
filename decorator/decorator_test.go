package decorator

import "testing"

func TestBasicWebPage_Display(t *testing.T) {
	website := NewBasicWebPage()
	website = NewAuthentication(website)
	website.Display()
}
