package main

import "fmt"

func main() {
	//Int
	var intdata int
	intdata = 10
	fmt.Println(intdata)

	var floatdata float32
	floatdata = 10.45
	fmt.Println(floatdata)

	var complexdata = complex(1, 2)
	fmt.Println(complexdata)

	var booldata bool
	booldata = true
	fmt.Println(booldata)

	var stringdata string
	stringdata = "Karthik"
	fmt.Println(stringdata)

	const constantdata = 10
	fmt.Println(constantdata)

	undeclared := 20
	fmt.Println(undeclared)

}
