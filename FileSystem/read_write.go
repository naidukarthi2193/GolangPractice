package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read File
	data, err := ioutil.ReadFile("./read.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	// Write File
	writedata := []byte("This is New Write Text")
	err2 := ioutil.WriteFile("write.txt", writedata, 0777)
	if err2 != nil {
		fmt.Println(err2)
	}

	//Append File
	f, err := os.OpenFile("write.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\nNew Sentence Added"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("write.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
