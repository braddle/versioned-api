package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	r.HandleFunc("/person", personCreateV1Handler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/vnd.person.v1+json")
	r.HandleFunc("/person", personCreateV1Handler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.ListenAndServe(":8080", r)
}

func personCreateV1Handler(w http.ResponseWriter, r *http.Request) {
	p1 := rest.PersonV1{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p1)

	p := person.Person{
		FirstName: p1.FirstName,
		LastName:  p1.LastName,
		Age:       p1.Age,
		HasTattoo: p1.HasTattoo,
	}

	person.Save(&p)

	w.WriteHeader(http.StatusCreated)
	n := negotiator.New(
		&rest.PersonV1Processor{},
		&rest.PersonV2Processor{},
		&rest.PersonV3Processor{},
	)
	if err := n.Negotiate(w, r, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
