package core

type Unstructured[ChildT Object] struct {
	Child Provider[ChildT]
}

func(unstruct Unstructured[ChildT]) Register(
	engine *Engine,
	namespace string,
) (children []Provider[ChildT], err RegistrationError) {
	if unstruct.Child != nil && unstruct.Child.IsStatic() {
		var child ChildT
		child, err = unstruct.Child.Provide(nil)
		if err == nil {
			err = child.Register(engine, namespace)
		}
	}
	if err != nil {
		children = []Provider[ChildT] {unstruct.Child}
	}
	return
}

func(unstruct Unstructured[ChildT]) Bind(*Engine) ([]Provider[ChildT], error) {
	return nil, nil
}

var _ Structured[Object] = Unstructured[Object]{}
