package core

import (
	"errors"
	"strings"
)

var listValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_LIST,
	name: "list",
}

type ListValue struct {
	Location *Location
	Value []Value
}

func(value *ListValue) Type() ValueType {
	return VT_LIST
}

func(value *ListValue) ComputationLocation() *Location {
	return value.Location
}

func(value *ListValue) Format(builder *strings.Builder) {
	builder.WriteRune('[')
	for index, child := range value.Value {
		if index > 0 {
			builder.WriteString(", ")
		}
		if child == nil {
			builder.WriteString("<nil value>")
		} else {
			child.Format(builder)
		}
	}
	builder.WriteRune(']')
}

func(value *ListValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, listValueTypeInfo, target)
}

func(value *ListValue) Iterator() Iterator {
	var elements []Value
	for _, child := range value.Value {
		if child != nil {
			elements = append(elements, child)
		}
	}
	return &SnapshotIterator {
		Values: elements,
	}
}

func ConvertListToBool(listValue Value) (boolValue Value, err error) {
	listInstance, ok := listValue.(*ListValue)
	if !ok {
		err = errors.New("Not given a *ListValue")
	} else {
		boolValue = &BoolValue {
			Location: listInstance.Location,
			Value: len(listInstance.Value) > 0,
		}
	}
	return
}

var _ IterableValue = &ListValue{}
