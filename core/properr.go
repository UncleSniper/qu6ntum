package core

import (
	"strings"
)

type PropertyValueConversionError struct {
	Location *Location
	ExpectedType string
	ConversionError error
}

func(err *PropertyValueConversionError) Error() string {
	var builder strings.Builder
	builder.WriteString("Cannot use value")
	if len(err.ExpectedType) > 0 {
		builder.WriteString(" as ")
		builder.WriteString(err.ExpectedType)
	}
	if err.Location != nil {
		builder.WriteString(" at ")
		builder.WriteString(err.Location.Format())
	}
	if err.ConversionError != nil {
		builder.WriteString(": ")
		builder.WriteString(err.ConversionError.Error())
	}
	return builder.String()
}

func(err *PropertyValueConversionError) IsWaitingForBind() bool {
	return false
}

var _ RegistrationError = &PropertyValueConversionError{}
