package core

import (
	"strings"
)

type IllegalMapKeyError struct {
	Key Value
}

func(err *IllegalMapKeyError) Error() string {
	var builder strings.Builder
	builder.WriteString("Illegal key for map value: ")
	if err.Key == nil {
		builder.WriteString("<nil value>")
		return builder.String()
	}
	keyType := err.Key.Type()
	switch keyType {
		case VT_INT:
			builder.WriteString("Int value with nil implementation")
			return builder.String()
		case VT_STRING:
			builder.WriteString("String value with nil implementation")
			return builder.String()
		case VT_OBJECT:
			objValue := err.Key.(*ObjectValue)
			if objValue == nil {
				builder.WriteString("Obj6ct value with nil implementation")
			} else if objValue.Value == nil {
				builder.WriteString("Obj6ct value with nil obj6ct")
			} else {
				builder.WriteString("Obj6ct value without OID")
			}
			return builder.String()
	}
	info := GetValueTypeInfo(keyType)
	var typeName string
	if info != nil {
		typeName = info.Name()
	}
	builder.WriteString("Value of ")
	if len(typeName) == 0 {
		builder.WriteString("unknown type")
	} else {
		builder.WriteString("type ")
		builder.WriteString(typeName)
	}
	return builder.String()
}
