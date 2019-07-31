package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV1Processor struct{}

func (p *PersonV1Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v1+json") ||
		strings.EqualFold(mediaRange, "application/json")
}

func (p *PersonV1Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v1+json")

	person, _ := dataModel.(person.Person)

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
