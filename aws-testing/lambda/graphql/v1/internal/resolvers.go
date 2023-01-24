package internal

import (
	"github.com/graph-gophers/graphql-go"
)

type inputQueryPerson struct {
	ID graphql.ID
}

type Queries struct {
	repository *personRepository
}

func (r *Queries) Person(input inputQueryPerson) *person {
	return r.repository.FindById(input.ID)
}

func (r *Queries) Hello() string {
	return "Hello World!"
}

func NewQueries(repository *personRepository) *Queries {
	return &Queries{
		repository: repository,
	}
}
