package core

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
)

var fileSizeValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_FILE_SIZE,
	name: "fileSize",
}

type FileSizeValue struct {
	Location *Location
	Value uint64
}

func(value *FileSizeValue) Type() ValueType {
	return VT_FILE_SIZE
}

func(value *FileSizeValue) ComputationLocation() *Location {
	return value.Location
}

func(value *FileSizeValue) Format(builder *strings.Builder) {
	builder.WriteString(strconv.FormatUint(value.Value, 10))
}

func(value *FileSizeValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, fileSizeValueTypeInfo, target)
}

func ConvertFileSizeToString(sizeValue Value) (strValue Value, err error) {
	sizeInstance, ok := sizeValue.(*FileSizeValue)
	if !ok {
		err = errors.New("Not given a *FileSizeValue")
	} else {
		strValue = &StringValue {
			Location: sizeInstance.Location,
			Value: strconv.FormatUint(sizeInstance.Value, 10),
		}
	}
	return
}

func ConvertFileSizeToBool(sizeValue Value) (boolValue Value, err error) {
	sizeInstance, ok := sizeValue.(*FileSizeValue)
	if !ok {
		err = errors.New("Not given a *FileSizeValue")
	} else {
		boolValue = &BoolValue {
			Location: sizeInstance.Location,
			Value: sizeInstance.Value != 0,
		}
	}
	return
}

func ConvertFileSizeToInt(sizeValue Value) (intValue Value, err error) {
	sizeInstance, ok := sizeValue.(*FileSizeValue)
	if !ok {
		err = errors.New("Not given a *FileSizeValue")
	} else {
		asInt := int64(sizeInstance.Value)
		if asInt < 0 {
			err = errors.New(fmt.Sprintf("Size %d exceeds range of an int", sizeInstance.Value))
		} else {
			intValue = &IntValue {
				Location: sizeInstance.Location,
				Value: asInt,
			}
		}
	}
	return
}

var _ Value = &FileSizeValue{}
