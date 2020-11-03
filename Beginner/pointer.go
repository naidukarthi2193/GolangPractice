package main

import "fmt"

func add(a *int) {
	fmt.Println(a)
	*a = *a + 2
}

func main() {
	var age int
	age = 11
	add(&age)
	fmt.Println(age)
}
