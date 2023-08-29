package core

type Structured[ChildT Object] interface {
	Register(*Engine, string) ([]Provider[ChildT], error)
	Bind(*Engine) ([]Provider[ChildT], error)
}
