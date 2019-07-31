package person

func GetById(id int) Person {
	return Person{
		Id:        id,
		FirstName: "Mark",
		LastName:  "Bradley",
		Age:       21,
		HasTattoo: false,
	}
}
