package core

type Provider[SubjectT any] interface {
	ProvisionLocation() *Location
	IsStatic() bool
	Bind(*Engine) error
	Provide(Context) (SubjectT, RegistrationError)
}
