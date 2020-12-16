package main

import (
	"fmt"
)

func myVariadicFunction(args ...string) {
	fmt.Println(args)
}

func myVariadicFunction2(args ...interface{}) {
	fmt.Println(args)
}

func main() {
	myVariadicFunction("hello", "world")

	myVariadicFunction2("hello", 2, true)
}
