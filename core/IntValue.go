package core

import (
	"fmt"
	"errors"
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

func ConvertIntToString(intValue Value) (strValue Value, err error) {
	intInstance, ok := intValue.(*IntValue)
	if !ok {
		err = errors.New("Not given a *IntValue")
	} else {
		strValue = &StringValue {
			Location: intInstance.Location,
			Value: strconv.FormatInt(intInstance.Value, 10),
		}
	}
	return
}

func ConvertIntToBool(intValue Value) (boolValue Value, err error) {
	intInstance, ok := intValue.(*IntValue)
	if !ok {
		err = errors.New("Not given a *IntValue")
	} else {
		boolValue = &BoolValue {
			Location: intInstance.Location,
			Value: intInstance.Value != 0,
		}
	}
	return
}

func ConvertIntToFileSize(intValue Value) (sizeValue Value, err error) {
	intInstance, ok := intValue.(*IntValue)
	if !ok {
		err = errors.New("Not given a *IntValue")
	} else if intInstance.Value < 0 {
		err = errors.New(fmt.Sprintf("Integer %d exceeds range of a fileSize", intInstance.Value))
	} else {
		sizeValue = &FileSizeValue {
			Location: intInstance.Location,
			Value: uint64(intInstance.Value),
		}
	}
	return
}

var _ Value = &IntValue{}
