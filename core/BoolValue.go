package core

import (
	"strings"
)

var boolValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_BOOL,
	name: "bool",
}

type BoolValue struct {
	Location *Location
	Value bool
}

func(value *BoolValue) Type() ValueType {
	return VT_BOOL
}

func(value *BoolValue) ComputationLocation() *Location {
	return value.Location
}

func(value *BoolValue) Format(builder *strings.Builder) {
	if value.Value {
		builder.WriteString("true")
	} else {
		builder.WriteString("false")
	}
}

func(value *BoolValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, boolValueTypeInfo, target)
}

var _ Value = &BoolValue{}
