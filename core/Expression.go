package core

type Expression interface {
	RenditionLocation() *Location
	Bind(*Engine) error
	Eval(Context) (Value, RegistrationError)
}
