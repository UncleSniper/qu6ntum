package core

import (
	"strings"
)

type ProviderAwaitingBindInBindError[SubjectT any] struct {
	Provider Provider[SubjectT]
	ClaimingError RegistrationError
}

func(err *ProviderAwaitingBindInBindError[SubjectT]) Error() string {
	var builder strings.Builder
	builder.WriteString("Provider")
	var location *Location
	if err.Provider != nil {
		location = err.Provider.ProvisionLocation()
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

var _ error = &ProviderAwaitingBindInBindError[int]{}
