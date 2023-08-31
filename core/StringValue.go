package core

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
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

func ConvertStringToInt(strValue Value) (intValue Value, err error) {
	strInstance, ok := strValue.(*StringValue)
	if !ok {
		err = errors.New("Not given a *StringValue")
		return
	}
	var theInt int64
	theInt, err = strconv.ParseInt(strInstance.Value, 0, 64)
	if err != nil {
		intValue = &IntValue {
			Location: strInstance.Location,
			Value: theInt,
		}
	}
	return
}

func ConvertStringToFloat(strValue Value) (floatValue Value, err error) {
	strInstance, ok := strValue.(*StringValue)
	if !ok {
		err = errors.New("Not given a *StringValue")
		return
	}
	var theFloat float64
	theFloat, err = strconv.ParseFloat(strInstance.Value, 64)
	if err != nil {
		floatValue = &FloatValue {
			Location: strInstance.Location,
			Value: theFloat,
		}
	}
	return
}

func ConvertStringToBool(strValue Value) (boolValue Value, err error) {
	strInstance, ok := strValue.(*StringValue)
	if !ok {
		err = errors.New("Not given a *StringValue")
	} else {
		str := strInstance.Value
		var theBool bool
		switch strings.ToLower(strings.TrimSpace(str)) {
			case "true", "on", "yes", "1", "enabled":
				theBool = true
			default:
				theBool = len(strInstance.Value) > 0
		}
		boolValue = &BoolValue {
			Location: strInstance.Location,
			Value: theBool,
		}
	}
	return
}

func ConvertStringToFileSize(strValue Value) (sizeValue Value, err error) {
	strInstance, ok := strValue.(*StringValue)
	if !ok {
		err = errors.New("Not given a *StringValue")
		return
	}
	var theSize uint64
	theSize, err = strconv.ParseUint(strInstance.Value, 0, 64)
	if err != nil {
		sizeValue = &FileSizeValue {
			Location: strInstance.Location,
			Value: theSize,
		}
	}
	return
}

var _ Value = &StringValue{}
