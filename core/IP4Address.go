package core

import (
	"net"
	"strings"
)

var ip4AddressTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_IP4_ADDRESS,
	name: "ip4Address",
}

type IP4Address struct {
	Location *Location
	Value net.IP
}

func(value *IP4Address) Type() ValueType {
	return VT_IP4_ADDRESS
}

func(value *IP4Address) ComputationLocation() *Location {
	return value.Location
}

func(value *IP4Address) Format(builder *strings.Builder) {
	if value.Value == nil {
		builder.WriteString("<nil IPv4 address>")
	} else {
		builder.WriteString(value.Value.String())
	}
}

func(value *IP4Address) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, ip4AddressTypeInfo, target)
}

var _ Value = &IP4Address{}
