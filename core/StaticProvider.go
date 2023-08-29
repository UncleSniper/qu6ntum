package core

type StaticProvider[SubjectT any] struct {
	Location *Location
	Subject SubjectT
}

func(provider *StaticProvider[SubjectT]) ProvisionLocation() *Location {
	return provider.Location
}

func(provider *StaticProvider[SubjectT]) IsStatic() bool {
	return true
}

func(provider *StaticProvider[SubjectT]) Bind(*Engine) error {
	return nil
}

func(provider *StaticProvider[SubjectT]) Provide(Context) (SubjectT, RegistrationError) {
	return provider.Subject, nil
}

var _ Provider[int] = &StaticProvider[int]{}
