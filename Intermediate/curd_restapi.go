package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Contact ...
type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Number int    `json:"number"`
}

var contacts []Contact

//HomePage ...
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("HomePage HIT")
}

//ReturnAllContacts ...
func ReturnAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
	fmt.Println("ALL endpoint HIT")

}

//ReturnContactID ...
func ReturnContactID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchAPI HIT")
	vars := mux.Vars(r)
	id := vars["id"]
	id1, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	for _, contact := range contacts {
		if contact.ID == id1 {
			json.NewEncoder(w).Encode(contact)
		}
	}
}

//CreateContact ...
func CreateContact(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Println(string(reqBody))
	var newContact Contact
	json.Unmarshal(reqBody, &newContact)
	contacts = append(contacts, newContact)
	json.NewEncoder(w).Encode(newContact)
}

//DeleteContactID ...
func DeleteContactID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	id1, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	for i, contact := range contacts {
		if contact.ID == id1 {
			contacts = append(contacts[:i], contacts[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(contacts)

}

func handlerequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/all", ReturnAllContacts)
	myRouter.HandleFunc("/contact/{id}", ReturnContactID)
	myRouter.HandleFunc("/create", CreateContact).Methods("POST")
	myRouter.HandleFunc("/delete/{id}", DeleteContactID).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func init() {
	contacts = []Contact{
		{1, "Karthikraj", "naidukarthi2193@gmail.com", 9773092096},
		{2, "Naidu", "email@gmail.com", 98216316320},
		{3, "Raj", "tests@gmail.com", 9705076543},
	}
}

func main() {
	handlerequests()
}
