package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	name string
	age  int
	id   int
}

func (employee *Employee) setname(name string) {
	employee.name = name
}
func (employee *Employee) setage(age int) {
	employee.age = age
}

func add(n1 int, n2 int) (string, error) {
	ans := n1 + n2
	final := strconv.Itoa(ans)
	return final, nil
}

func subtract() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}

}

func main() {
	fmt.Println("Hello ")
	var employee Employee
	employee.setname("karthikraj")
	employee.setage(1)
	fmt.Println(employee)

	ans, err := add(5, 6)
	if err == nil {
		fmt.Println(ans)
	}

	fmt.Println("kasr")
}
