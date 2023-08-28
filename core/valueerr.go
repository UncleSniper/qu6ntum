package core

import (
	"strings"
)

type ValueConversionError struct {
	SourceValue Value
	SourceType *ValueTypeInfo
	TargetType ValueType
	ConversionError error
}

func(err *ValueConversionError) Error() string {
	var builder strings.Builder
	builder.WriteString("Failed to convert ")
	if err.SourceValue == nil {
		builder.WriteString("nil value")
	} else {
		sourceInfo := err.SourceType
		if sourceInfo == nil {
			sourceInfo = GetValueTypeInfo(err.SourceType.Type())
		}
		if sourceInfo != nil {
			typeName := sourceInfo.Name()
			if len(typeName) == 0 {
				typeName = "<anonymous type>"
			}
			builder.WriteString(typeName)
			builder.WriteRune(' ')
		}
		err.SourceValue.Format(&builder)
	}
	targetInfo := GetValueTypeInfo(err.TargetType)
	if targetInfo == nil {
		builder.WriteString(" to unknown target type")
	} else {
		builder.WriteString(" to ")
		typeName := targetInfo.Name()
		if len(typeName) == 0 {
			typeName = "<anonymous type>"
		}
		builder.WriteString(typeName)
	}
	if err.ConversionError == nil {
		builder.WriteString(": No known conversion")
	} else {
		builder.WriteString(": ")
		builder.WriteString(err.ConversionError.Error())
	}
	return builder.String()
}

type ValueConverterReturnedWrongTypeError struct {
	TargetType ValueType
	ReturnedValue Value
}

func(err *ValueConverterReturnedWrongTypeError) Error() string {
	if err.ReturnedValue == nil {
		return "Converter returned nil value"
	}
	var builder strings.Builder
	builder.WriteString("Converter returned value of ")
	info := GetValueTypeInfo(err.ReturnedValue.Type())
	if info == nil {
		builder.WriteString("unknown claimed type")
	} else {
		builder.WriteString("claimed type ")
		typeName := info.Name()
		if len(typeName) == 0 {
			typeName = "<anonymous type>"
		}
		builder.WriteString(typeName)
	}
	builder.WriteString(", which does not actually satisfy target type")
	return builder.String()
}
