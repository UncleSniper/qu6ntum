package core

type Structured[ChildT Object] interface {
	Register(*Engine, string) ([]Provider[ChildT], RegistrationError)
	Bind(*Engine) ([]Provider[ChildT], error)
}
