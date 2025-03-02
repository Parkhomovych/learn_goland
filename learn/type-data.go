package learn

import (
	"fmt"
	"reflect"
)

func TypeData() {
	var name string = "Vlad"
	var age int = 25
	var isBoolean bool = true
	var balance float64 = 1000.0

	fmt.Println("string", name)
	fmt.Println("int", age)
	fmt.Println("bool", isBoolean)
	fmt.Println("float", balance)
	fmt.Println("type", reflect.TypeOf(name))
}
