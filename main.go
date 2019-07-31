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
	r.HandleFunc("/person", personCreateV2Handler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/vnd.person.v2+json")
	r.HandleFunc("/person", personCreateV3Handler).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/vnd.person.v3+json")
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

	renderPerson(w, r, p, http.StatusCreated)
}

func personCreateV2Handler(w http.ResponseWriter, r *http.Request) {
	p2 := rest.PersonV2{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p2)

	p := person.Person{
		FirstName: p2.FirstName,
		LastName:  p2.LastName,
		HasTattoo: p2.HasTattoo,
	}

	person.Save(&p)

	renderPerson(w, r, p, http.StatusCreated)
}

func personCreateV3Handler(w http.ResponseWriter, r *http.Request) {
	p3 := rest.PersonV3{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p3)

	p := person.Person{
		FirstName:   p3.FirstName,
		LastName:    p3.LastName,
		HasTattoo:   p3.HasTattoo,
		HasPiercing: p3.HasPiercing,
	}

	person.Save(&p)

	renderPerson(w, r, p, http.StatusCreated)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	p := person.GetById(id)

	renderPerson(w, r, p, http.StatusOK)
}

func renderPerson(w http.ResponseWriter, r *http.Request, p person.Person, statusCode int) {
	w.WriteHeader(statusCode)
	n := negotiator.New(
		&rest.PersonV1Processor{},
		&rest.PersonV2Processor{},
		&rest.PersonV3Processor{},
	)
	if err := n.Negotiate(w, r, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
