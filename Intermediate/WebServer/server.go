package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

//IncrementCounter ...
func IncrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	//Serve Static Files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//Increment Counter
	http.HandleFunc("/add", IncrementCounter)

	//View Counter
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", counter)
	})

	//Serve
	log.Fatal(http.ListenAndServeTLS(":8000", "server.crt", "server.key", nil))
	/*
		HTTPS Certificate
		openssl genrsa -out server.key 2048
		openssl ecparam -genkey -name secp384r1 -out server.key
		openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	*/
}
