package core

type VariableDefinition struct {
	DefinitionLocation *Location
	NameProvider Provider[string]
	Namespace string
	RenderedName string
	QualifiedName string
	nameDetermined bool
	nameWaitingForBindError RegistrationError
	ValueProvider Provider[Value]
	InitialValue Value
	valueDetermined bool
	Variable *Variable
}

func(definition *VariableDefinition) Location() *Location {
	return definition.DefinitionLocation
}

func(definition *VariableDefinition) Register(engine *Engine, namespace string) RegistrationError {
	if !definition.nameDetermined {
		definition.Namespace = namespace
		if definition.NameProvider == nil {
			return &VariableMissingNameError {
				Definition: definition,
			}
		}
		err := definition.NameProvider.Bind(engine)
		if err != nil {
			return &FailedToDetermineVariableNameError {
				Definition: definition,
				DeterminationError: err,
			}
		}
		providedName, regErr := definition.NameProvider.Provide(nil)
		if regErr != nil {
			if regErr.IsWaitingForBind() {
				return nil
			}
			return regErr
		}
		definition.RenderedName = providedName
		definition.QualifiedName = JoinNames(namespace, providedName)
		definition.nameDetermined = true
		if definition.Variable == nil {
			definition.Variable = &Variable {
				Location: definition.DefinitionLocation,
				Name: definition.QualifiedName,
			}
			if !engine.SetVariable(definition.QualifiedName, definition.Variable) {
				return &VariableNameClashError {
					Definition: definition,
					PreviousVariable: engine.GetVariable(definition.QualifiedName),
				}
			}
		}
	}
	if !definition.valueDetermined && definition.Variable != nil {
		if definition.ValueProvider == nil {
			definition.Variable.Value = nil
		} else {
			err := definition.ValueProvider.Bind(engine)
			if err != nil {
				return &FailedToDetermineVariableValueError {
					Definition: definition,
					DeterminationError: err,
				}
			}
			providedValue, regErr := definition.ValueProvider.Provide(nil)
			if regErr != nil {
				if regErr.IsWaitingForBind() {
					return nil
				}
				return regErr
			}
			definition.InitialValue = providedValue
			definition.valueDetermined = true
			if definition.Variable.Value == nil {
				definition.Variable.Value = providedValue
			}
		}
	}
	return nil
}

func(definition *VariableDefinition) Bind(engine *Engine) error {
	if !definition.nameDetermined {
		if definition.NameProvider == nil {
			return &VariableMissingNameError {
				Definition: definition,
			}
		}
		err := definition.NameProvider.Bind(engine)
		if err != nil {
			return &FailedToDetermineVariableNameError {
				Definition: definition,
				DeterminationError: err,
			}
		}
		providedName, regErr := definition.NameProvider.Provide(nil)
		if regErr != nil {
			return nil
		}
		definition.RenderedName = providedName
		definition.QualifiedName = JoinNames(definition.Namespace, providedName)
		definition.nameDetermined = true
		if definition.Variable == nil {
			definition.Variable = &Variable {
				Location: definition.DefinitionLocation,
				Name: definition.QualifiedName,
			}
			if !engine.SetVariable(definition.QualifiedName, definition.Variable) {
				return &VariableNameClashError {
					Definition: definition,
					PreviousVariable: engine.GetVariable(definition.QualifiedName),
				}
			}
		}
	}
	if !definition.valueDetermined && definition.Variable != nil {
		if definition.ValueProvider == nil {
			definition.Variable.Value = nil
		} else {
			err := definition.ValueProvider.Bind(engine)
			if err != nil {
				return &FailedToDetermineVariableValueError {
					Definition: definition,
					DeterminationError: err,
				}
			}
			providedValue, regErr := definition.ValueProvider.Provide(nil)
			if regErr != nil {
				return regErr
			}
			definition.InitialValue = providedValue
			definition.valueDetermined = true
			if definition.Variable.Value == nil {
				definition.Variable.Value = providedValue
			}
		}
	}
	return nil
}

var _ Definition = &VariableDefinition{}
