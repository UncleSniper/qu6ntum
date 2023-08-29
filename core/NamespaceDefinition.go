package core

type NamespaceDefinition struct {
	DefinedAt *Location
	Name string
	Children []Definition
}

func(ns *NamespaceDefinition) Location() *Location {
	return ns.DefinedAt
}

func(ns *NamespaceDefinition) Register(engine *Engine, outerNamespace string) (err RegistrationError) {
	innerNamespace := JoinNames(outerNamespace, ns.Name)
	for _, child := range ns.Children {
		err = child.Register(engine, innerNamespace)
		if err != nil {
			return
		}
	}
	return
}

func(ns *NamespaceDefinition) Bind(engine *Engine) (err error) {
	for _, child := range ns.Children {
		err = child.Bind(engine)
		if err != nil {
			return
		}
	}
	return
}

var _ Definition = &NamespaceDefinition{}
