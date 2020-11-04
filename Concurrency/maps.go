package main

import "fmt"

//Service ...
type Service interface {
	Hello()
}

// MyService ...
type MyService struct{}

// Hello ...
func (service MyService) Hello() {
	fmt.Println("Hello 1")
}

// MyService2 ...
type MyService2 struct{}

// Hello ....
func (service MyService2) Hello() {
	fmt.Println("Hello 2")
}

func main() {
	mymap := make(map[string]int)
	mymap["karthik"] = 1
	mymap["kar"] = 1
	mymap["thik"] = 1
	mymap["karik"] = 1
	for key, val := range mymap {
		fmt.Println(key, val)
	}
	fmt.Println(mymap)
	delete(mymap, "kar")
	fmt.Println(mymap)

	interfacemaps := make(map[string]Service)
	interfacemaps["1"] = MyService{}
	interfacemaps["2"] = MyService2{}

	for _, val := range interfacemaps {
		val.Hello()
	}

}
