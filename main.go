package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/braddle/versioned-api/person"
	"github.com/braddle/versioned-api/rest"

	"github.com/gorilla/mux"
	"github.com/jchannon/negotiator"
)

func main() {
	fmt.Println("Staring Server")
	r := mux.NewRouter()
	r.HandleFunc("/person/{id}", personHandler).Methods(http.MethodGet)
	r.HandleFunc("/person", personCreateHandler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/vnd.person.v1+json")
	r.HandleFunc("/person", personCreateHandler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.ListenAndServe(":8080", r)
}

func personCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	p := person.GetById(id)

	n := negotiator.New(
		&rest.PersonV1Processor{},
		&rest.PersonV2Processor{},
		&rest.PersonV3Processor{},
	)
	if err := n.Negotiate(w, r, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
