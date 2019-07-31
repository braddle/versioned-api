package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jchannon/negotiator"
)

type person struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	HasTattoo bool
}

func main() {
	fmt.Println("Staring Server")
	r := mux.NewRouter()
	r.HandleFunc("/person/{id}", personHandler)

	http.ListenAndServe(":8080", r)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	p := person{
		Id:        id,
		FirstName: "Mark",
		LastName:  "Bradley",
		Age:       21,
		HasTattoo: false,
	}

	if err := negotiator.Negotiate(w, r, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
