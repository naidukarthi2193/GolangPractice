package main

import (
	"fmt"
)

var score int

func init() {
	fmt.Println("init function started")
	fmt.Println(score)
	score = 1
}

func main() {

	fmt.Println("import main functions")
	fmt.Println(score)
}
