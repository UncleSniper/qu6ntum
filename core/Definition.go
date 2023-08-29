package core

type Definition interface {
	Location() *Location
	Register(*Engine, string) RegistrationError
	Bind(*Engine) error
}

type RegistrationError interface {
	error
	IsWaitingForBind() bool
}
