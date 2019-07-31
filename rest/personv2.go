package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV2Processor struct{}

func (p *PersonV2Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v2+json")
}

func (p *PersonV2Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v1+json")

	person, _ := dataModel.(person.Person)

	p1 := map[string]interface{}{
		"id":         person.Id,
		"first_name": person.FirstName,
		"last_name":  person.LastName,
		"has_tattoo": person.HasTattoo,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
