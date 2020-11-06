package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Temp File Create
	f, err := ioutil.TempFile("", "temp")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	//Temp File write
	if _, err := f.Write([]byte("hello world\n")); err != nil {
		fmt.Println(err)
	}

	//Temp File Read
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))

	// Temp Dir
	tempDir, err := ioutil.TempDir("", "cars-")
	if err != nil {
		fmt.Println(err)
	}
	defer os.RemoveAll(tempDir)

	// Temp File Create
	file, err := ioutil.TempFile(tempDir, "car-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())

	if _, err := file.Write([]byte("hello world\n")); err != nil {
		fmt.Println(err)
	}

	data1, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data1))
}
