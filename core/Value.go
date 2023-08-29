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

func identityConversion(value Value) (Value, error) {
	return value, nil
}

var knownValueTypes []*ValueTypeInfo

func NewValueType(name string) *ValueTypeInfo {
	if knownValueTypes == nil {
		knownValueTypes = []*ValueTypeInfo {
			intValueTypeInfo,
			floatValueTypeInfo,
			stringValueTypeInfo,
			boolValueTypeInfo,
			objectValueTypeInfo,
			listValueTypeInfo,
			mapValueTypeInfo,
			ip4AddressTypeInfo,
			ip6AddressTypeInfo,
			ip4SocketAddressTypeInfo,
			ip6SocketAddressTypeInfo,
			iterableValueTypeInfo,
		}
	}
	if len(name) == 0 {
		name = "<anonymous type>"
	}
	info := &ValueTypeInfo {
		typeID: ValueType(len(knownValueTypes) + 1),
		name: name,
		conversions: make(map[ValueType]ValueConversion),
	}
	info.conversions[info.typeID] = identityConversion
	knownValueTypes = append(knownValueTypes, info)
	return info
}

func GetValueTypeInfo(id ValueType) *ValueTypeInfo {
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
