package decorator

type WebPageDecorator struct {
	wrappee WebSite
}

func NewWenPageDecorator(wrappee WebSite) WebSite {
	return WebPageDecorator{
		wrappee: wrappee,
	}
}

func (wpd WebPageDecorator) Display() {
	wpd.wrappee.Display()
}

func (wpd WebPageDecorator) Close() string {
	return wpd.wrappee.Close()
}
