package core

type Structured[ChildT Object] interface {
	Register(*Engine, string) ([]ChildT, error)
}
