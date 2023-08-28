package core

import (
	"fmt"
	"strings"
	"strconv"
)

var objectValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_OBJECT,
	name: "obj6ct",
}

type ObjectValue struct {
	Location *Location
	Value Object
}

func(value *ObjectValue) Type() ValueType {
	return VT_OBJECT
}

func(value *ObjectValue) ComputationLocation() *Location {
	return value.Location
}

func FormatObjectValue(object Object, builder *strings.Builder) {
	if object == nil {
		builder.WriteString("<nil obj6ct>")
		return
	}
	builder.WriteString("<obj6ct")
	id := object.ID()
	if id != NO_OID {
		builder.WriteString(" #")
		builder.WriteString(strconv.FormatUint(uint64(id), 10))
	}
	implType := object.ImplementationType()
	if len(implType) > 0 {
		builder.WriteString(" of implementation type ")
		builder.WriteString(implType)
	}
	descr := object.Description()
	if len(descr) > 0 {
		builder.WriteString(fmt.Sprintf(", description = %q", descr))
	}
	builder.WriteRune('>')
}

func(value *ObjectValue) Format(builder *strings.Builder) {
	FormatObjectValue(value.Value, builder)
}

func(value *ObjectValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, objectValueTypeInfo, target)
}

var _ Value = &ObjectValue{}
