package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV3Processor struct{}

type PersonV3 struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	HasTattoo   bool   `json:"has_tattoo"`
	HasPiercing bool   `json:"has_piercing"`
}

func (p *PersonV3Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v3+json")
}

func (p *PersonV3Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v3+json")

	person, _ := dataModel.(person.Person)

	p1 := PersonV3{
		Id:          person.Id,
		FirstName:   person.FirstName,
		LastName:    person.LastName,
		HasTattoo:   person.HasTattoo,
		HasPiercing: person.HasPiercing,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
