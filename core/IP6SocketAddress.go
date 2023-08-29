package core

import (
	"net"
	"strings"
	"strconv"
)

var ip6SocketAddressTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_IP6_SOCKET_ADDRESS,
	name: "ip6SocketAddress",
}

type IP6SocketAddress struct {
	Location *Location
	IP net.IPAddr
	Port uint16
}

func(value *IP6SocketAddress) Type() ValueType {
	return VT_IP6_SOCKET_ADDRESS
}

func(value *IP6SocketAddress) ComputationLocation() *Location {
	return value.Location
}

func(value *IP6SocketAddress) Format(builder *strings.Builder) {
	if value.IP.IP == nil || value.Port == 0 {
		builder.WriteString("<nil IPv6 socket address>")
	} else {
		builder.WriteString(value.IP.String())
		builder.WriteRune(':')
		builder.WriteString(strconv.FormatUint(uint64(value.Port), 10))
	}
}

func(value *IP6SocketAddress) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, ip6SocketAddressTypeInfo, target)
}

var _ Value = &IP6SocketAddress{}
