package core

type Unstructured[ChildT Object] struct {
	Child ChildT
}

func(unstruct Unstructured[ChildT]) Register(engine *Engine, namespace string) (children []ChildT, err error) {
	err = unstruct.Child.Register(engine, namespace)
	if err != nil {
		children = []ChildT {unstruct.Child}
	}
	return
}

var _ Structured[Object] = Unstructured[Object]{}
