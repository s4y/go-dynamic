package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

type Potato interface {
	Peel()
}

type Russet struct{}

func (p *Russet) Peel() {}

type derp struct{}

type Notato Russet

func TestMakeType(t *testing.T) {
	russetType := reflect.TypeOf(Russet{})
	notatoType := reflect.TypeOf(Notato{})
	dynamicPotatoType := MakeType("DynamicPotato", reflect.TypeOf(struct{}{}))

	AddMethod(reflect.PtrTo(dynamicPotatoType), "Peel", func(p *struct{}) {
		fmt.Println("Made it!")
	})

	fmt.Printf("Potato: %s %#v\n", russetType.String(), fromReflectType(russetType).methods)
	fmt.Printf("Notato: %s %#v\n", notatoType.String(), fromReflectType(notatoType).methods)
	fmt.Printf("DynamicPotato: %s %#v\n", dynamicPotatoType.String(), fromReflectType(dynamicPotatoType))

	dynamicPotato, ok := reflect.ValueOf(&struct{}{}).Convert(reflect.PtrTo(dynamicPotatoType)).Interface().(Potato)
	if !ok {
		t.Error("Our fresh type isn't convertible to Potato after adding the right methods")
	}

	var potato Potato = dynamicPotato
	potato = &Russet{}

	potato.Peel()

	// (*reflect.Value)(fromReflectType(reflect.PtrTo(dynamicPotatoType)).methods[0].ifn).Call([]reflect.Value{
	// 	reflect.Indirect(reflect.ValueOf(dynamicPotato)).Addr().Convert(reflect.PtrTo(reflect.TypeOf(struct{}{}))),
	// })

	dynamicPotato.Peel()
}
