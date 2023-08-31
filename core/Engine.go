package core

import (
	"sync"
)

type Engine struct {
	objects map[string]Object
	variables map[string]*Variable
	stateLock sync.Mutex
}

func(engine *Engine) SetObject(name string, object Object) bool {
	engine.stateLock.Lock()
	if engine.objects == nil {
		if object == nil {
			engine.stateLock.Unlock()
			return false
		}
		engine.objects = make(map[string]Object)
	}
	have := engine.objects[name] != nil
	if object == nil {
		if !have {
			engine.stateLock.Unlock()
			return false
		}
		delete(engine.objects, name)
	} else {
		if have {
			engine.stateLock.Unlock()
			return false
		}
		engine.objects[name] = object
	}
	engine.stateLock.Unlock()
	return true
}

func(engine *Engine) GetObject(name string) (object Object) {
	engine.stateLock.Lock()
	if engine.objects != nil {
		object = engine.objects[name]
	}
	engine.stateLock.Unlock()
	return
}

func(engine *Engine) SetVariable(name string, variable *Variable) bool {
	engine.stateLock.Lock()
	if engine.variables == nil {
		if variable == nil {
			engine.stateLock.Unlock()
			return false
		}
		engine.variables = make(map[string]*Variable)
	}
	have := engine.variables[name] != nil
	if variable == nil {
		if !have {
			engine.stateLock.Unlock()
			return false
		}
		delete(engine.variables, name)
	} else {
		if have {
			engine.stateLock.Unlock()
			return false
		}
		engine.variables[name] = variable
	}
	engine.stateLock.Unlock()
	return true
}

func(engine *Engine) GetVariable(name string) (variable *Variable) {
	engine.stateLock.Lock()
	if engine.variables != nil {
		variable = engine.variables[name]
	}
	engine.stateLock.Unlock()
	return
}
