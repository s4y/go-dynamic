package dynamic

import (
	"fmt"
	"reflect"
	"unsafe"
)

func trampoline() {
	fmt.Println("BOUNCE BOUNCE")
}

func AddMethod(t reflect.Type, name string, f interface{}) {
	fv := reflect.ValueOf(f)
	ft := fv.Type()
	if ft.NumIn() == 0 {
		panic("AddMethod: function must take at least one parameter (the receiver)")
	}
	if !ft.In(0).ConvertibleTo(t) {
		panic("AddMethod: function's first parameter must be convertible to " + t.String())
	}
	inTypes := make([]reflect.Type, ft.NumIn())
	for i := 0; i < len(inTypes); i++ {
		inTypes[i] = ft.In(i)
	}
	outTypes := make([]reflect.Type, ft.NumOut())
	for i := 0; i < len(outTypes); i++ {
		outTypes[i] = ft.Out(i)
	}

	tfn := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		fmt.Println("TFN!!!!!!!!!!!!!!!!")
		return fv.Call(args)
	})
	_ = reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		fmt.Println("IFN!!!!!!!!!!!!!!!!")
		return fv.Call(append([]reflect.Value{reflect.Indirect(args[0])}, args[1:]...))
	})

	rt := fromReflectType(t)
	rt.methods = append(rt.methods, method{
		name: &name,
		mtyp: fromReflectType(reflect.FuncOf(inTypes[1:], outTypes, ft.IsVariadic())),
		typ:  fromReflectType(reflect.FuncOf(append([]reflect.Type{t}, inTypes...), outTypes, ft.IsVariadic())),
		ifn:  unsafe.Pointer(&tfn),
		tfn:  unsafe.Pointer(&tfn),
	})
}
