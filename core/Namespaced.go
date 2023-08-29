package core

type Namespaced[ChildT Object] struct {
	Name string
	Children []Structured[ChildT]
}

func(ns *Namespaced[ChildT]) Register(engine *Engine, outerNamespace string) ([]Provider[ChildT], error) {
	innerNamespace := JoinNames(outerNamespace, ns.Name)
	var children []Provider[ChildT]
	for _, ref := range ns.Children {
		grandchildren, err := ref.Register(engine, innerNamespace)
		if err != nil {
			return nil, err
		}
		children = append(children, grandchildren...)
	}
	return children, nil
}

func(ns *Namespaced[ChildT]) Bind(engine *Engine) ([]Provider[ChildT], error) {
	var children []Provider[ChildT]
	for _, ref := range ns.Children {
		grandchildren, err := ref.Bind(engine)
		if err != nil {
			return nil, err
		}
		children = append(children, grandchildren...)
	}
	return children, nil
}

func JoinNames(names ...string) string {
	var all []byte
	for _, name := range names {
		seenSlash := true
		for index := 0; index < len(name); index++ {
			c := name[index]
			if c == '/' {
				seenSlash = true
				if index == 0 {
					all = nil
				}
			} else {
				if seenSlash {
					if len(all) > 0 {
						all = append(all, '/')
					}
					seenSlash = false
				}
				all = append(all, c)
			}
		}
	}
	return string(all)
}

var _ Structured[Object] = &Namespaced[Object]{}
