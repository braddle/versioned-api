package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV1Processor struct{}

type PersonV1 struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasTattoo bool   `json:"has_tattoo"`
}

func (p *PersonV1Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v1+json") ||
		strings.EqualFold(mediaRange, "application/json")
}

func (p *PersonV1Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v1+json")

	person, _ := dataModel.(person.Person)

	p1 := PersonV1{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
		HasTattoo: person.HasTattoo,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
