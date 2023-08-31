package core

type Property[TargetT any] interface {
	Register(*Engine, string, TargetT) RegistrationError
	Bind(*Engine, TargetT) error
}

type MappedProperty[OuterTargetT any, InnerTargetT any] struct {
	Property Property[InnerTargetT]
	MapTarget func(OuterTargetT) (InnerTargetT, RegistrationError)
}

func(prop *MappedProperty[OuterTargetT, InnerTargetT]) Register(
	engine *Engine,
	namespace string,
	outer OuterTargetT,
) RegistrationError {
	if prop.Property == nil || prop.MapTarget == nil {
		return nil
	}
	inner, err := prop.MapTarget(outer)
	if err != nil {
		return err
	}
	return prop.Property.Register(engine, namespace, inner)
}

func(prop *MappedProperty[OuterTargetT, InnerTargetT]) Bind(engine *Engine, outer OuterTargetT) error {
	if prop.Property == nil || prop.MapTarget == nil {
		return nil
	}
	inner, err := prop.MapTarget(outer)
	if err != nil {
		return err
	}
	return prop.Property.Bind(engine, inner)
}

const (
	ifPropFl_HAVE_CONDITION int = 1 << iota
	ifPropFl_SATISFIED
)

type IfProperty[TargetT any] struct {
	Location *Location
	Condition Provider[bool]
	Then Property[TargetT]
	Else Property[TargetT]
	Namespace string
	flags int
}

func(prop *IfProperty[TargetT]) evalCondition() RegistrationError {
	if (prop.flags & ifPropFl_HAVE_CONDITION) != 0 {
		return nil
	}
	if prop.Condition == nil {
		prop.flags |= ifPropFl_HAVE_CONDITION
		return nil
	}
	value, err := prop.Condition.Provide(nil)
	if err != nil {
		return err
	}
	prop.flags |= ifPropFl_HAVE_CONDITION
	if value {
		prop.flags |= ifPropFl_SATISFIED
	}
	return nil
}

func(prop *IfProperty[TargetT]) getBranch() Property[TargetT] {
	if (prop.flags & ifPropFl_SATISFIED) != 0 {
		return prop.Then
	} else {
		return prop.Else
	}
}

func(prop *IfProperty[TargetT]) Register(
	engine *Engine,
	namespace string,
	target TargetT,
) (err RegistrationError) {
	prop.Namespace = namespace
	err = prop.evalCondition()
	if err != nil {
		if err.IsWaitingForBind() {
			return nil
		}
		return err
	}
	branch := prop.getBranch()
	if branch == nil {
		return nil
	}
	return branch.Register(engine, namespace, target)
}

func(prop *IfProperty[TargetT]) Bind(engine *Engine, target TargetT) (err error) {
	hadCondition := (prop.flags & ifPropFl_HAVE_CONDITION) != 0
	if !hadCondition {
		evalErr := prop.evalCondition()
		if evalErr != nil {
			if evalErr.IsWaitingForBind() {
				err = &ProviderAwaitingBindInBindError[bool] {
					Provider: prop.Condition,
					ClaimingError: evalErr,
				}
			} else {
				err = evalErr
			}
			return
		}
	}
	branch := prop.getBranch()
	if branch == nil {
		return nil
	}
	if !hadCondition {
		err = branch.Register(engine, prop.Namespace, target)
		if err != nil {
			return
		}
	}
	err = branch.Bind(engine, target)
	return
}

var _ Property[int] = &MappedProperty[int, bool]{}
var _ Property[int] = &IfProperty[int]{}
