package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Users ...
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

// User ...
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// Social ...
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	file, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}
	var users Users
	xmlfile, err := ioutil.ReadAll(file)
	xml.Unmarshal(xmlfile, &users)
	fmt.Println(users)
}
