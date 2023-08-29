package core

type ExpressionProvider[SubjectT any] struct {
	Location *Location
	Expression Expression
	Conversion func(Value) (SubjectT, RegistrationError)
}

func(provider *ExpressionProvider[SubjectT]) ProvisionLocation() *Location {
	if provider.Location != nil {
		return provider.Location
	}
	if provider.Expression == nil {
		return nil
	}
	return provider.Expression.RenditionLocation()
}

func(provider *ExpressionProvider[SubjectT]) IsStatic() bool {
	return false
}

func(provider *ExpressionProvider[SubjectT]) Bind(engine *Engine) error {
	if provider.Expression == nil {
		return nil
	}
	return provider.Expression.Bind(engine)
}

func(provider *ExpressionProvider[SubjectT]) Provide(context Context) (subject SubjectT, err RegistrationError) {
	if provider.Expression == nil {
		return
	}
	var value Value
	value, err = provider.Expression.Eval(context)
	if err == nil && provider.Conversion != nil {
		subject, err = provider.Conversion(value)
	}
	return
}

var _ Provider[int] = &ExpressionProvider[int]{}
