package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/braddle/versioned-api/person"
)

type PersonV3Processor struct{}

func (p *PersonV3Processor) CanProcess(mediaRange string) bool {
	return strings.EqualFold(mediaRange, "application/vnd.person.v3+json")
}

func (p *PersonV3Processor) Process(w http.ResponseWriter, req *http.Request, dataModel interface{}, context ...interface{}) error {
	w.Header().Set("Content-Type", "application/vnd.person.v3+json")

	person, _ := dataModel.(person.Person)

	p1 := map[string]interface{}{
		"id":           person.Id,
		"first_name":   person.FirstName,
		"last_name":    person.LastName,
		"has_tattoo":   person.HasTattoo,
		"has_piercing": person.HasPiercing,
	}

	j, _ := json.Marshal(p1)

	w.Write(j)

	return nil
}
