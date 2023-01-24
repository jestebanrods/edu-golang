package internal

import "github.com/graph-gophers/graphql-go"

type personRepository struct {
	data map[graphql.ID]*person
}

func (r *personRepository) FindById(id graphql.ID) *person {
	return r.data[id]
}

func NewPersonRepositoy() *personRepository {
	people := []*person{
		{
			ID:        "1000",
			FirstName: "Pedro",
			LastName:  "Marquez",
		},

		{
			ID:        "1001",
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	data := make(map[graphql.ID]*person, len(people))

	for _, p := range people {
		data[p.ID] = p
	}

	return &personRepository{
		data: data,
	}
}
