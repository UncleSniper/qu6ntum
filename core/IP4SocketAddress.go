package core

import (
	"net"
	"strings"
	"strconv"
)

var ip4SocketAddressTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_IP4_SOCKET_ADDRESS,
	name: "ip4SocketAddress",
}

type IP4SocketAddress struct {
	Location *Location
	IP net.IP
	Port uint16
}

func(value *IP4SocketAddress) Type() ValueType {
	return VT_IP4_SOCKET_ADDRESS
}

func(value *IP4SocketAddress) ComputationLocation() *Location {
	return value.Location
}

func(value *IP4SocketAddress) Format(builder *strings.Builder) {
	if value.IP == nil || value.Port == 0 {
		builder.WriteString("<nil IPv4 socket address>")
	} else {
		builder.WriteString(value.IP.String())
		builder.WriteRune(':')
		builder.WriteString(strconv.FormatUint(uint64(value.Port), 10))
	}
}

func(value *IP4SocketAddress) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, ip4SocketAddressTypeInfo, target)
}

var _ Value = &IP4SocketAddress{}
