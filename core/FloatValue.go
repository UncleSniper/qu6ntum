package core

import (
	"fmt"
	"strings"
)

var floatValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_FLOAT,
	name: "float",
}

type FloatValue struct {
	Location *Location
	Value float64
}

func(value *FloatValue) Type() ValueType {
	return VT_FLOAT
}

func(value *FloatValue) ComputationLocation() *Location {
	return value.Location
}

func(value *FloatValue) Format(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("%G", value.Value))
}

func(value *FloatValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, floatValueTypeInfo, target)
}

var _ Value = &FloatValue{}
