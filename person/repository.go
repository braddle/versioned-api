package person

import "math/rand"

func GetById(id int) Person {
	return Person{
		Id:          id,
		FirstName:   "Mark",
		LastName:    "Bradley",
		Age:         21,
		HasTattoo:   false,
		HasPiercing: false,
	}
}

func Save(p *Person) {
	p.Id = rand.Int()
}
