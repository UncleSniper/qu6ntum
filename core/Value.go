package core

import (
	"strings"
)

type ValueType uint

const (
	VT_INT ValueType = iota + 1
	VT_FLOAT
	VT_STRING
	VT_BOOL
	VT_FILE_SIZE
	VT_OBJECT
	VT_LIST
	VT_MAP
	VT_IP4_ADDRESS
	VT_IP6_ADDRESS
	VT_IP4_SOCKET_ADDRESS
	VT_IP6_SOCKET_ADDRESS
	VT_ITERABLE
)

type ValueTypeInfo struct {
	typeID ValueType
	name string
	conversions map[ValueType]ValueConversion
}

type ValueConversion func(Value) (Value, error)

func(info *ValueTypeInfo) Type() ValueType {
	return info.typeID
}

func(info *ValueTypeInfo) Name() string {
	return info.name
}

func(info *ValueTypeInfo) ConversionTo(target ValueType) ValueConversion {
	if info.conversions == nil {
		return nil
	} else {
		return info.conversions[target]
	}
}

func(info *ValueTypeInfo) SetConversion(target ValueType, conversion ValueConversion) {
	if target == 0 {
		return
	}
	if info.conversions == nil {
		info.conversions = make(map[ValueType]ValueConversion)
	}
	info.conversions[target] = conversion
}

func IdentityConversion(value Value) (Value, error) {
	return value, nil
}

var knownValueTypes []*ValueTypeInfo

func initKnownValueTypes() {
	knownValueTypes = []*ValueTypeInfo {
		intValueTypeInfo,
		floatValueTypeInfo,
		stringValueTypeInfo,
		boolValueTypeInfo,
		fileSizeValueTypeInfo,
		objectValueTypeInfo,
		listValueTypeInfo,
		mapValueTypeInfo,
		ip4AddressTypeInfo,
		ip6AddressTypeInfo,
		ip4SocketAddressTypeInfo,
		ip6SocketAddressTypeInfo,
		iterableValueTypeInfo,
	}
	intValueTypeInfo.SetConversion(VT_INT, IdentityConversion)
	intValueTypeInfo.SetConversion(VT_STRING, ConvertIntToString)
	intValueTypeInfo.SetConversion(VT_BOOL, ConvertIntToBool)
	intValueTypeInfo.SetConversion(VT_FILE_SIZE, ConvertIntToFileSize)
	floatValueTypeInfo.SetConversion(VT_FLOAT, IdentityConversion)
	stringValueTypeInfo.SetConversion(VT_STRING, IdentityConversion)
	stringValueTypeInfo.SetConversion(VT_INT, ConvertStringToInt)
	stringValueTypeInfo.SetConversion(VT_FLOAT, ConvertStringToFloat)
	stringValueTypeInfo.SetConversion(VT_BOOL, ConvertStringToBool)
	stringValueTypeInfo.SetConversion(VT_FILE_SIZE, ConvertStringToFileSize)
	boolValueTypeInfo.SetConversion(VT_BOOL, IdentityConversion)
	fileSizeValueTypeInfo.SetConversion(VT_FILE_SIZE, IdentityConversion)
	fileSizeValueTypeInfo.SetConversion(VT_STRING, ConvertFileSizeToString)
	fileSizeValueTypeInfo.SetConversion(VT_BOOL, ConvertFileSizeToBool)
	fileSizeValueTypeInfo.SetConversion(VT_INT, ConvertFileSizeToInt)
	objectValueTypeInfo.SetConversion(VT_OBJECT, IdentityConversion)
	listValueTypeInfo.SetConversion(VT_LIST, IdentityConversion)
	listValueTypeInfo.SetConversion(VT_BOOL, ConvertListToBool)
	listValueTypeInfo.SetConversion(VT_ITERABLE, IdentityConversion)
	mapValueTypeInfo.SetConversion(VT_MAP, IdentityConversion)
	mapValueTypeInfo.SetConversion(VT_BOOL, ConvertMapToBool)
	mapValueTypeInfo.SetConversion(VT_ITERABLE, IdentityConversion)
	ip4AddressTypeInfo.SetConversion(VT_IP4_ADDRESS, IdentityConversion)
	ip4SocketAddressTypeInfo.SetConversion(VT_IP4_SOCKET_ADDRESS, IdentityConversion)
	ip6AddressTypeInfo.SetConversion(VT_IP6_ADDRESS, IdentityConversion)
	ip6SocketAddressTypeInfo.SetConversion(VT_IP6_SOCKET_ADDRESS, IdentityConversion)
	iterableValueTypeInfo.SetConversion(VT_ITERABLE, IdentityConversion)
}

func NewValueType(name string) *ValueTypeInfo {
	if knownValueTypes == nil {
		initKnownValueTypes()
	}
	if len(name) == 0 {
		name = "<anonymous type>"
	}
	info := &ValueTypeInfo {
		typeID: ValueType(len(knownValueTypes) + 1),
		name: name,
		conversions: make(map[ValueType]ValueConversion),
	}
	info.conversions[info.typeID] = IdentityConversion
	knownValueTypes = append(knownValueTypes, info)
	return info
}

func GetValueTypeInfo(id ValueType) *ValueTypeInfo {
	if knownValueTypes == nil {
		initKnownValueTypes()
	}
	if id == 0 || int(id) > len(knownValueTypes) {
		return nil
	}
	return knownValueTypes[int(id) - 1]
}

type Value interface {
	Type() ValueType
	ComputationLocation() *Location
	Format(*strings.Builder)
	GetAs(ValueType) (Value, error)
}

func ValueAsString(value Value) string {
	if value == nil {
		return "<nil value>"
	}
	var builder strings.Builder
	value.Format(&builder)
	return builder.String()
}

func ValueAsStringInto(value Value, builder *strings.Builder) {
	if value == nil {
		builder.WriteString("<nil value>")
	} else {
		value.Format(builder)
	}
}

func GetTypedValueAs(value Value, info *ValueTypeInfo, target ValueType) (newValue Value, err error) {
	if value == nil {
		err = &ValueConversionError {
			TargetType: target,
		}
		return
	}
	if info == nil {
		info = GetValueTypeInfo(target)
	}
	var conversion ValueConversion
	if info != nil {
		if info.Type() == target {
			newValue = value
			return
		}
		conversion = info.ConversionTo(target)
	}
	if conversion == nil {
		err = &ValueConversionError {
			SourceValue: value,
			SourceType: info,
			TargetType: target,
		}
		return
	}
	newValue, err = conversion(value)
	if err != nil {
		err = &ValueConversionError {
			SourceValue: value,
			SourceType: info,
			TargetType: target,
			ConversionError: err,
		}
	}
	return
}

func GetValueAs[TargetT Value](value Value, target ValueType) (newValue TargetT, err error) {
	if value == nil {
		err = &ValueConversionError {
			TargetType: target,
		}
		return
	}
	var returned Value
	returned, err = value.GetAs(target)
	if err != nil {
		return
	}
	var ok bool
	newValue, ok = returned.(TargetT)
	if !ok {
		err = &ValueConversionError {
			SourceValue: value,
			SourceType: GetValueTypeInfo(value.Type()),
			TargetType: target,
			ConversionError: &ValueConverterReturnedWrongTypeError {
				TargetType: target,
				ReturnedValue: returned,
			},
		}
	}
	return
}
