package dynamic

import (
	"reflect"
	"unsafe"
)

func AddMethod(t reflect.Type, name string, f interface{}) {
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	if ft.NumIn() == 0 {
		panic("AddMethod: function must take at least one parameter (the receiver)")
	}
	if !ft.In(0).ConvertibleTo(t) {
		panic("AddMethod: function's first parameter must be convertible to " + t.String())
	}
	inTypes := make([]reflect.Type, ft.NumIn()-1)
	for i := 0; i < len(inTypes); i++ {
		inTypes[i] = ft.In(i + 1)
	}
	outTypes := make([]reflect.Type, ft.NumOut())
	for i := 0; i < len(outTypes); i++ {
		outTypes[i] = ft.Out(i)
	}

	tfn := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		return fv.Call(args)
	})
	// TODO: Check for kindDirectIface
	ifn := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		return fv.Call(append([]reflect.Value{reflect.Indirect(args[0])}, args[1:]...))
	})

	rt := fromReflectType(t)
	rt.methods = append(rt.methods, method{
		name: &name,
		mtyp: fromReflectType(reflect.FuncOf(inTypes, outTypes, ft.IsVariadic())),
		typ:  fromReflectType(reflect.FuncOf(append([]reflect.Type{t}, inTypes...), outTypes, ft.IsVariadic())),
		ifn:  unsafe.Pointer(&ifn),
		tfn:  unsafe.Pointer(&tfn),
	})
}
