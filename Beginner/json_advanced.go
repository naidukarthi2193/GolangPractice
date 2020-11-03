package main

import (
	"encoding/json"
	"fmt"
)

//Person ...
type Person struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Account interface{} `json:"account"`
}

//Account ...
type Account struct {
	Number int  `json:"number"`
	Member bool `json:"member"`
}

func main() {

	peeps := []Person{
		{"Karthik", 2, nil},
		{"Karthik", 2, Account{1, true}},
		{"Karthik", 2, Account{2, true}},
		{"Karthik", 2, nil},
		{"Karthik", 2, Account{3, false}},
	}
	bytearr, err := json.MarshalIndent(peeps, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytearr))

	jsonstr := string(bytearr)
	var reader []map[string]interface{}
	err = json.Unmarshal([]byte(jsonstr), &reader)
	for _, a := range reader {
		fmt.Printf("%+v\n", a)
	}
}
