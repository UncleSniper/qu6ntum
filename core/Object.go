package core

import (
	"sync/atomic"
)

type OID uint64

const NO_OID OID = 0

type Object interface {
	ID() OID
	ImplementationType() string
	DefinitionLocation() *Location
	Description() string
	Register(*Engine, string) error
}

var nextOID atomic.Uint64

func NextOID() OID {
	return OID(nextOID.Add(1))
}
