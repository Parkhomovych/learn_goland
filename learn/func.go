package learn

import "fmt"

func MainFunc() {
	withParams("Vlad")
	fmt.Println(withParamsAndReturn("Vlad"))
}

func withParams(name string) {
	fmt.Println("Hello, " + name + "!")
}

func withParamsAndReturn(name string) string {
	return "Hello, " + name + "!"
}
