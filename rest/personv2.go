package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV2Processor struct{}

type PersonV2 struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasTattoo bool   `json:"has_tattoo"`
}

func (p *PersonV2Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v2+json")
}

func (p *PersonV2Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v2+json")

	person, _ := dataModel.(person.Person)

	p1 := PersonV2{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		HasTattoo: person.HasTattoo,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
