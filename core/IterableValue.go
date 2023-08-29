package core

var iterableValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_ITERABLE,
	name: "iterable",
}

type IterableValue interface {
	Value
	Iterator() Iterator
}

type Iterator interface {
	Next() (Value, bool)
}

type SnapshotIterator struct {
	Values []Value
	nextIndex int
}

func(iterator *SnapshotIterator) Next() (value Value, valid bool) {
	length := len(iterator.Values)
	if iterator.nextIndex >= length {
		return
	}
	value = iterator.Values[iterator.nextIndex]
	iterator.nextIndex++
	valid = true
	return
}

var _ Iterator = &SnapshotIterator{}
