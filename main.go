package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	n := negotiator.New(&personV1Processor{})
	if err := n.Negotiate(w, r, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type personV1Processor struct{}

func (p *personV1Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v1+json") ||
		strings.EqualFold(mediaRange, "application/json")
}

func (p *personV1Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v1+json")

	person, _ := dataModel.(person)

	p1 := map[string]interface{}{
		"id":         person.Id,
		"first_name": person.FirstName,
		"last_name":  person.LastName,
		"age":        person.Age,
		"has_tattoo": person.HasTattoo,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
