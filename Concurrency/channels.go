package main

import (
	"fmt"
	"time"
)

//Send ...
func Send(c chan string) {
	fmt.Println("Start")
	time.Sleep(1 * time.Second)
	c <- "hello"
	fmt.Println("Send")
}

func main() {
	fmt.Println("Chanel")
	values := make(chan string, 2)
	defer close(values)

	go Send(values)
	go Send(values)

	val := <-values
	fmt.Println(val)
	time.Sleep(1 * time.Second)

}
