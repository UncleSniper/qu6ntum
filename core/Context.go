package core

import (
	"sync/atomic"
)

type CID uint64

const NO_CID CID = 0

type Context interface {
	ID() CID
}

var nextCID atomic.Uint64

func NextCID() CID {
	return CID(nextCID.Add(1))
}
