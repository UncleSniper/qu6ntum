package core

import (
	"fmt"
	"strings"
	"strconv"
)

var mapValueTypeInfo *ValueTypeInfo = &ValueTypeInfo {
	typeID: VT_MAP,
	name: "map",
}

type MapValue struct {
	Location *Location
	ints map[int64]Value
	strings map[string]Value
	objects map[OID]mapPair[Object]
}

type mapPair[KeyT any] struct {
	key KeyT
	value Value
}

func(value *MapValue) Type() ValueType {
	return VT_MAP
}

func(value *MapValue) ComputationLocation() *Location {
	return value.Location
}

func(value *MapValue) Format(builder *strings.Builder) {
	builder.WriteRune('{')
	var had bool
	if value.ints != nil {
		for key, val := range value.ints {
			if had {
				builder.WriteString(", ")
			} else {
				had = true
			}
			builder.WriteString(strconv.FormatInt(key, 10))
			builder.WriteString(": ")
			if val == nil {
				builder.WriteString("<nil value>")
			} else {
				val.Format(builder)
			}
		}
	}
	if value.strings != nil {
		for key, val := range value.strings {
			if had {
				builder.WriteString(", ")
			} else {
				had = true
			}
			builder.WriteString(fmt.Sprintf("%q", key))
			builder.WriteString(": ")
			if val == nil {
				builder.WriteString("<nil value>")
			} else {
				val.Format(builder)
			}
		}
	}
	if value.objects != nil {
		for _, pair := range value.objects {
			if had {
				builder.WriteString(", ")
			} else {
				had = true
			}
			FormatObjectValue(pair.key, builder)
			builder.WriteString(": ")
			if pair.value == nil {
				builder.WriteString("<nil value>")
			} else {
				pair.value.Format(builder)
			}
		}
	}
	builder.WriteRune('}')
}

func(value *MapValue) GetAs(target ValueType) (Value, error) {
	return GetTypedValueAs(value, mapValueTypeInfo, target)
}

func(theMap *MapValue) PutInt(key int64, value Value) {
	if theMap.ints == nil {
		if value == nil {
			return
		}
		theMap.ints = make(map[int64]Value)
	}
	if value == nil {
		delete(theMap.ints, key)
	} else {
		theMap.ints[key] = value
	}
}

func(theMap *MapValue) GetInt(key int64) Value {
	if theMap.ints == nil {
		return nil
	}
	return theMap.ints[key]
}

func(theMap *MapValue) PutString(key string, value Value) {
	if theMap.strings == nil {
		if value == nil {
			return
		}
		theMap.strings = make(map[string]Value)
	}
	if value == nil {
		delete(theMap.strings, key)
	} else {
		theMap.strings[key] = value
	}
}

func(theMap *MapValue) GetString(key string) Value {
	if theMap.strings == nil {
		return nil
	}
	return theMap.strings[key]
}

func(theMap *MapValue) PutObject(key Object, value Value) error {
	var id OID
	if key != nil {
		id = key.ID()
	}
	if id == NO_OID {
		return &IllegalMapKeyError {
			Key: &ObjectValue {
				Value: key,
			},
		}
	}
	theMap.putObject(key, id, value)
	return nil
}

func(theMap *MapValue) putObject(key Object, id OID, value Value) {
	if theMap.objects == nil {
		if value == nil {
			return
		}
		theMap.objects = make(map[OID]mapPair[Object])
	}
	if value == nil {
		delete(theMap.objects, id)
	} else {
		theMap.objects[id] = mapPair[Object] {
			key: key,
			value: value,
		}
	}
}

func(theMap *MapValue) GetObject(key Object) Value {
	var id OID
	if key != nil {
		id = key.ID()
	}
	if id == NO_OID || theMap.objects == nil {
		return nil
	}
	return theMap.objects[id].value
}

func(theMap *MapValue) Put(key Value, value Value) error {
	if key == nil {
		return &IllegalMapKeyError{}
	}
	switch key.Type() {
		case VT_INT:
			intKey := key.(*IntValue)
			if intKey != nil {
				theMap.PutInt(intKey.Value, value)
				return nil
			}
		case VT_STRING:
			stringKey := key.(*StringValue)
			if stringKey != nil {
				theMap.PutString(stringKey.Value, value)
				return nil
			}
		case VT_OBJECT:
			objectKey := key.(*ObjectValue)
			if objectKey != nil {
				var id OID
				if objectKey.Value != nil {
					id = objectKey.Value.ID()
				}
				if id != NO_OID {
					theMap.putObject(objectKey.Value, id, value)
					return nil
				}
			}
	}
	return &IllegalMapKeyError {
		Key: key,
	}
}

func(theMap *MapValue) Get(key Value) Value {
	if key == nil {
		return nil
	}
	switch key.Type() {
		case VT_INT:
			intKey := key.(*IntValue)
			if intKey != nil {
				return theMap.GetInt(intKey.Value)
			}
		case VT_STRING:
			stringKey := key.(*StringValue)
			if stringKey != nil {
				return theMap.GetString(stringKey.Value)
			}
		case VT_OBJECT:
			objectKey := key.(*ObjectValue)
			if objectKey != nil {
				return theMap.GetObject(objectKey.Value)
			}
	}
	return nil
}

func(theMap *MapValue) Size() int {
	var size int
	if theMap.ints != nil {
		size = len(theMap.ints)
	}
	if theMap.strings != nil {
		size += len(theMap.strings)
	}
	if theMap.objects != nil {
		size += len(theMap.objects)
	}
	return size
}

var _ Value = &MapValue{}
