package core

import (
	"strings"
)

type ExpressionAwaitingBindInBindError struct {
	Expression Expression
	ClaimingError RegistrationError
}

func(err *ExpressionAwaitingBindInBindError) Error() string {
	var builder strings.Builder
	builder.WriteString("Expression")
	var location *Location
	if err.Expression != nil {
		location = err.Expression.RenditionLocation()
	}
	if location != nil {
		builder.WriteString(" rendered at ")
		builder.WriteString(location.Format())
	}
	builder.WriteString(" claims to await bind, but we're already in bind phase")
	if err.ClaimingError != nil {
		builder.WriteString(", claiming error is: ")
		builder.WriteString(err.ClaimingError.Error())
	}
	return builder.String()
}

var _ error = &ExpressionAwaitingBindInBindError{}
