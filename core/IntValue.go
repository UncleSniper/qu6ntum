package core

import (
	"strings"
	"strconv"
)

var intValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_INT,
	name: "int",
}

type IntValue struct {
	Location *Location
	Value int64
}

func(value *IntValue) Type() ValueType {
	return VT_INT
}

func(value *IntValue) ComputationLocation() *Location {
	return value.Location
}

func(value *IntValue) Format(builder *strings.Builder) {
	builder.WriteString(strconv.FormatInt(value.Value, 10))
}

func(value *IntValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, intValueTypeInfo, target)
}

var _ Value = &IntValue{}
