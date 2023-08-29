package core

type Definition interface {
	Location() *Location
	Register(*Engine, string) error
	Bind(*Engine) error
}
