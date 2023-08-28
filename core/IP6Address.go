package core

import (
	"net"
	"strings"
)

var ip6AddressTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_IP6_ADDRESS,
	name: "ip6Address",
}

type IP6Address struct {
	Location *Location
	// Zone -> ALPHA / DIGIT / "-" / "." / "_" / "~"
	Value net.IPAddr
}

func(value *IP6Address) Type() ValueType {
	return VT_IP6_ADDRESS
}

func(value *IP6Address) ComputationLocation() *Location {
	return value.Location
}

func(value *IP6Address) Format(builder *strings.Builder) {
	if value.Value.IP == nil {
		builder.WriteString("nil IPv6 address>")
	} else {
		builder.WriteString(value.Value.String())
	}
}

func(value *IP6Address) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, ip6AddressTypeInfo, target)
}

var _ Value = &IP6Address{}
