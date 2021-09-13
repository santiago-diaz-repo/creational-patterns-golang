package chainOfResponsability

type SecProvider interface {
	Validate() bool
}
