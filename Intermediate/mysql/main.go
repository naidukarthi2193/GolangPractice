package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

//Users ...
type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "668lHo9aEw:OK5hzEt3JE@tcp(remotemysql.com:3306)/668lHo9aEw")
	if err != nil {
		panic(err)
	}
	return db

}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("ALlUsers")
	db = getDB()
	results, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(results)

	var alluser []Users
	for results.Next() {
		var users Users
		err = results.Scan(&users.ID, &users.Name, &users.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(users)
		alluser = append(alluser, users)
		log.Printf(users.Name)
	}
	results.Close()
	db.Close()

	json.NewEncoder(w).Encode(alluser)

}

func newUser(w http.ResponseWriter, r *http.Request) {
	db = getDB()
	fmt.Fprintf(w, "New User Endpoint Hit")
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	nuser := Users{rand.Intn(200), name, email}
	s := fmt.Sprintf("INSERT INTO test (id, name, email) VALUES ( %d, '%s', '%s')", nuser.ID, nuser.Name, nuser.Email)
	fmt.Println(s)
	insert, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
	db.Close()
	fmt.Fprintf(w, "Add User Endpoint Hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete User")
	db = getDB()
	name := mux.Vars(r)["name"]
	s := fmt.Sprintf("DELETE FROM test WHERE name='%s'", name)
	del, err := db.Query(s)
	if err != nil {
		panic(err)
	}
	del.Close()
	db.Close()
	fmt.Fprintf(w, "Deleted")

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating User")
	db = getDB()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	s := fmt.Sprintf("UPDATE test SET email='%s' WHERE name='%s'", email, name)
	updt, err := db.Query(s)
	if err != nil {
		panic(err)
	}
	updt.Close()
	db.Close()
	fmt.Fprintf(w, "Update User Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/delete/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/update/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/add/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	db, err := sql.Open("mysql", "668lHo9aEw:OK5hzEt3JE@tcp(remotemysql.com:3306)/668lHo9aEw")
	if err != nil {
		panic(err)
	}
	_, err1 := db.Query("DELETE FROM test")
	if err1 != nil {
		panic(err1.Error())
	}
	fmt.Println("Server Started")
	handleRequests()
	defer db.Close()

}
