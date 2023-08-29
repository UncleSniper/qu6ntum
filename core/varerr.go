package core

import (
	"strings"
)

type VariableMissingNameError struct {
	Definition *VariableDefinition
}

func(err *VariableMissingNameError) Error() string {
	var builder strings.Builder
	builder.WriteString("Variable definition")
	if err.Definition != nil && err.Definition.DefinitionLocation != nil {
		builder.WriteString(" at ")
		builder.WriteString(err.Definition.DefinitionLocation.Format())
	}
	builder.WriteString(" has no name Provider")
	return builder.String()
}

func(err *VariableMissingNameError) IsWaitingForBind() bool {
	return false
}

type FailedToDetermineVariableNameError struct {
	Definition *VariableDefinition
	DeterminationError error
}

func(err *FailedToDetermineVariableNameError) Error() string {
	var builder strings.Builder
	builder.WriteString("Failed to determine variable name for variable definition")
	if err.Definition != nil && err.Definition.DefinitionLocation != nil {
		builder.WriteString(" at ")
		builder.WriteString(err.Definition.DefinitionLocation.Format())
	}
	var message string
	if err.DeterminationError != nil {
		message = err.DeterminationError.Error()
	}
	if len(message) > 0 {
		builder.WriteString(": ")
		builder.WriteString(message)
	}
	return builder.String()
}

func(err *FailedToDetermineVariableNameError) IsWaitingForBind() bool {
	return false
}

type VariableNameClashError struct {
	Definition *VariableDefinition
	PreviousVariable *Variable
}

func(err *VariableNameClashError) Error() string {
	var builder strings.Builder
	builder.WriteString("Variable name clash")
	if err.Definition != nil {
		builder.WriteString(" for name '")
		builder.WriteString(err.Definition.QualifiedName)
		builder.WriteRune('\'')
	}
	if err.PreviousVariable != nil && err.PreviousVariable.Location != nil {
		builder.WriteString(" (previous definition at ")
		builder.WriteString(err.PreviousVariable.Location.Format())
		builder.WriteRune(')')
	}
	return builder.String()
}

func(err *VariableNameClashError) IsWaitingForBind() bool {
	return false
}

type FailedToDetermineVariableValueError struct {
	Definition *VariableDefinition
	DeterminationError error
}

func(err *FailedToDetermineVariableValueError) Error() string {
	var builder strings.Builder
	builder.WriteString("Failed to determine initial value for variable definition")
	if err.Definition != nil && err.Definition.DefinitionLocation != nil {
		builder.WriteString(" at ")
		builder.WriteString(err.Definition.DefinitionLocation.Format())
	}
	var message string
	if err.DeterminationError != nil {
		message = err.DeterminationError.Error()
	}
	if len(message) > 0 {
		builder.WriteString(": ")
		builder.WriteString(message)
	}
	return builder.String()
}

func(err *FailedToDetermineVariableValueError) IsWaitingForBind() bool {
	return false
}

var _ RegistrationError = &VariableMissingNameError{}
var _ RegistrationError = &FailedToDetermineVariableNameError{}
var _ RegistrationError = &VariableNameClashError{}
var _ RegistrationError = &FailedToDetermineVariableValueError{}
