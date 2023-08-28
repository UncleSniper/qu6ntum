package core

import (
	"fmt"
	"strconv"
)

type Location struct {
	File string
	Line uint
	Column uint
}

func(location *Location) isEmpty() bool {
	return len(location.File) == 0 && location.Line == 0
}

func(location *Location) Format() string {
	if location == nil || location.isEmpty() {
		return "<unknown location>"
	}
	var f, l, c string
	if len(location.File) == 0 {
		f = "<unknown file>"
	} else {
		f = location.File
	}
	if location.Line == 0 {
		l = "<unknown line>"
	} else {
		l = strconv.FormatUint(uint64(location.Line), 10)
	}
	if location.Column == 0 {
		c = ""
	} else {
		c = fmt.Sprintf(":%d", location.Column)
	}
	return fmt.Sprintf("%s:%s%s", f, l, c)
}
