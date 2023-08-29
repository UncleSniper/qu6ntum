package core

import (
	"net"
)

type ip4AddressParseState uint

type IP4AddressParser struct {
	octets []byte
	octetBytes uint
	currentOctet byte
	charIndex int
}

func(parser *IP4AddressParser) pushDigit(c byte) *ValueParseError {
	if c < byte('0') || c > byte('9') {
		var expected string
		if parser.octetBytes == 0 {
			expected = "digit"
		} else if parser.octetBytes == 3 {
			if len(parser.octets) >= 3 {
				expected = "end-of-address"
			} else {
				expected = "'.'"
			}
		} else if len(parser.octets) >= 3 {
			expected = "digit or end-of-address"
		} else {
			expected = "digit or '.'"
		}
		return &ValueParseError {
			Expected: expected,
			Found: c,
			HasFound: true,
			Index: parser.charIndex,
		}
	}
	var nextOctet byte = parser.currentOctet * 10 + (c - byte('0'))
	if nextOctet < parser.currentOctet {
		var expected string
		if len(parser.octets) >= 3 {
			expected = "end-of-address"
		} else {
			expected = "'.'"
		}
		return &ValueParseError {
			Expected: expected,
			Found: c,
			HasFound: true,
			Index: parser.charIndex,
		}
	}
	parser.currentOctet = nextOctet
	parser.octetBytes++
	parser.charIndex++
	return nil
}

func(parser *IP4AddressParser) PushByte(c byte) (err *ValueParseError) {
	if len(parser.octets) < 4 {
		if c == byte('.') {
			if parser.octetBytes == 0 {
				return &ValueParseError {
					Expected: "digit",
					Found: c,
					HasFound: true,
					Index: parser.charIndex,
				}
			}
			parser.octets = append(parser.octets, parser.currentOctet)
			parser.octetBytes = 0
			parser.currentOctet = 0
			parser.charIndex++
		} else {
			err = parser.pushDigit(c)
		}
	} else {
		err = &ValueParseError {
			Expected: "end-of-address",
			Found: c,
			HasFound: true,
			Index: parser.charIndex,
		}
	}
	return
}

func(parser *IP4AddressParser) End() (net.IP, *ValueParseError) {
	switch len(parser.octets) {
		case 3:
			if parser.octetBytes == 0 {
				return nil, &ValueParseError {
					Expected: "digit",
					Index: parser.charIndex,
				}
			}
			parser.octets = append(parser.octets, parser.currentOctet)
			parser.currentOctet = 0
			parser.octetBytes = 0
			return net.IP(parser.octets), nil
		case 4:
			return net.IP(parser.octets), nil
		default:
			var expected string
			if parser.octetBytes == 0 {
				expected = "digit"
			} else {
				expected = "digit or '.'"
			}
			return nil, &ValueParseError {
				Expected: expected,
				Index: parser.charIndex,
			}
	}
}
