package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

type Potato interface {
	Peel()
}

type RussetPotato struct{}

func (p *RussetPotato) Peel() {}

func TestMakeType(t *testing.T) {
	russetPotatoType := reflect.TypeOf(RussetPotato{})
	dynamicPotatoType := MakeType("DynamicPotato", reflect.TypeOf(struct{}{}))

	AddMethod(reflect.PtrTo(dynamicPotatoType), "Peel", func(p *struct{}) {
		fmt.Println("Made it!")
	})

	fmt.Printf("RussetPotato: %s %#v\n", russetPotatoType.String(), fromReflectType(russetPotatoType).ptrToThis.methods)
	fmt.Printf("DynamicPotato: %s %#v\n", dynamicPotatoType.String(), fromReflectType(dynamicPotatoType).ptrToThis.methods)

	dynamicPotato := reflect.ValueOf(&struct{}{}).Convert(reflect.PtrTo(dynamicPotatoType))

	// Segfault!
	dynamicPotato.MethodByName("Peel").Call([]reflect.Value{})

	potato, ok := dynamicPotato.Interface().(Potato)
	if !ok {
		t.Error("Our fresh type isn't convertible to Potato after adding the right methods")
	}

	// (*reflect.Value)(fromReflectType(reflect.PtrTo(dynamicPotatoType)).methods[0].ifn).Call([]reflect.Value{
	// 	reflect.Indirect(reflect.ValueOf(dynamicPotato)).Addr().Convert(reflect.PtrTo(reflect.TypeOf(struct{}{}))),
	// })

	potato.Peel()
}
