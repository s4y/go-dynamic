package dynamic

import (
	"reflect"
	"unsafe"
)

func copyType(t reflect.Type) *rtype {
	p := unsafe.Pointer(fromReflectType(t))
	if t.Kind() < reflect.Array {
		var ret rtype = *(*rtype)(p)
		return &ret
	}
	switch t.Kind() {
	case reflect.Array:
		ret := *(*arrayType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Chan:
		ret := *(*chanType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Func:
		ret := *(*funcType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Map:
		ret := *(*mapType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Ptr:
		ret := *(*ptrType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Slice:
		ret := *(*sliceType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	case reflect.Struct:
		ret := *(*structType)(p)
		return (*rtype)(unsafe.Pointer(&ret))
	default:
		panic("copyType called on unsupported type " + t.String())
	}
}

func MakeType(name string, underlyingType reflect.Type) reflect.Type {
	t := copyType(underlyingType)
	typeOut := toReflectType(t)
	t.string = &name
	t.uncommonType = &uncommonType{name: t.string}
	t.ptrToThis = nil
	t.ptrToThis = fromReflectType(reflect.PtrTo(typeOut))
	t.ptrToThis.uncommonType = &uncommonType{methods: []method{}}
	return typeOut
}
