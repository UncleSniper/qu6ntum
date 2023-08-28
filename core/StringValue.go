package core

import (
	"fmt"
	"strings"
)

var stringValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_STRING,
	name: "string",
}

type StringValue struct {
	Location *Location
	Value string
}

func(value *StringValue) Type() ValueType {
	return VT_STRING
}

func(value *StringValue) ComputationLocation() *Location {
	return value.Location
}

func(value *StringValue) Format(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("%q", value.Value))
}

func(value *StringValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, stringValueTypeInfo, target)
}

var _ Value = &StringValue{}
