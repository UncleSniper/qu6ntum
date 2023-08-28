package core

type NamespaceDefinition struct {
	DefinedAt *Location
	Name string
	Children []Definition
}

func(ns *NamespaceDefinition) Location() *Location {
	return ns.DefinedAt
}

func(ns *NamespaceDefinition) Define(engine *Engine, outerNamespace string) (err error) {
	innerNamespace := JoinNames(outerNamespace, ns.Name)
	for _, child := range ns.Children {
		err = child.Define(engine, innerNamespace)
		if err != nil {
			return
		}
	}
	return
}

var _ Definition = &NamespaceDefinition{}
