package dynamic

import (
	"reflect"
	"unsafe"
)

func fromReflectType(t reflect.Type) *rtype {
	return (*rtype)((*(*emptyInterface)(unsafe.Pointer(&t))).word)
}

func toReflectType(t *rtype) reflect.Type {
	var iptr interface{}
	(*emptyInterface)(unsafe.Pointer(&iptr)).typ = t
	return reflect.TypeOf(iptr)
}
