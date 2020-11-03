package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Users ...
type Users struct {
	Users []User `json:"users"`
}

// User ...
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

//Social ...
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	bytearray, err := ioutil.ReadAll(file)

	var users Users
	json.Unmarshal(bytearray, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println(users.Users[i])
	}

}
