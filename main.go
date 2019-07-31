package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	j, _ := json.Marshal(p)

	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}
