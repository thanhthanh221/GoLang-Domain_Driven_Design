package Aggregate

import (
	"GolangDDD/Entities"
	"GolangDDD/ValueObject"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidPersion = errors.New("a cutsomer has to have a valid name")
)

type Customer struct {
	person       *Entities.Person
	products     []*Entities.Item
	transactions []ValueObject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPersion
	}

	person := &Entities.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*Entities.Item, 0),
		transactions: make([]ValueObject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &Entities.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &Entities.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
