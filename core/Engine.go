package core

type Engine struct {
	objects map[string]Object
	variables map[string]*Variable
}

func(engine *Engine) SetObject(name string, object Object) bool {
	if engine.objects == nil {
		if object == nil {
			return false
		}
		engine.objects = make(map[string]Object)
	}
	have := engine.objects[name] != nil
	if object == nil {
		if !have {
			return false
		}
		delete(engine.objects, name)
	} else {
		if have {
			return false
		}
		engine.objects[name] = object
	}
	return true
}

func(engine *Engine) GetObject(name string) Object {
	if engine.objects == nil {
		return nil
	}
	return engine.objects[name]
}

func(engine *Engine) SetVariable(name string, variable *Variable) bool {
	if engine.variables == nil {
		if variable == nil {
			return false
		}
		engine.variables = make(map[string]*Variable)
	}
	have := engine.variables[name] != nil
	if variable == nil {
		if !have {
			return false
		}
		delete(engine.variables, name)
	} else {
		if have {
			return false
		}
		engine.variables[name] = variable
	}
	return true
}

func(engine *Engine) GetVariable(name string) *Variable {
	if engine.variables == nil {
		return nil
	}
	return engine.variables[name]
}
