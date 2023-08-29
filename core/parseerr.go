package core

import (
	"fmt"
	"strings"
)

type ValueParseError struct {
	Expected string
	HasFound bool
	Found byte
	Index int
}

func(err *ValueParseError) Error() string {
	var builder strings.Builder
	builder.WriteString("Expected ")
	builder.WriteString(err.Expected)
	if err.Index >= 0 {
		builder.WriteString(fmt.Sprintf(" at index %d", err.Index))
	}
	if err.HasFound {
		builder.WriteString(fmt.Sprintf(", but found %q", err.Found))
	} else {
		builder.WriteString(", but found end-of-input")
	}
	return builder.String()
}

var _ error = &ValueParseError {}
