package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	type Employee struct {
		id     int
		person Person
	}

	var arraydata []string
	arraydata = []string{"k", "a", "r", "t"}
	newdata := [...]string{"k", "a", "r", "t"}
	fmt.Println(arraydata)
	fmt.Println(newdata[1:3])

	namesage := map[string]int{
		"TutorialEdge":  2240,
		"MKBHD":         6580350,
		"Geeksforgeeks": 171220,
	}
	fmt.Println(namesage)

	newperson := Person{"karthikraj", 1}
	fmt.Println(newperson.age)
	newemployee1 := Employee{1, Person{"Naidu", 1}}
	fmt.Println(newemployee1)

}
