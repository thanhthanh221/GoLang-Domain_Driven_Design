package Memory

import (
	"GolangDDD/Aggregate"
	"GolangDDD/Domain/Customer"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	Customer map[uuid.UUID]Aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		Customer: make(map[uuid.UUID]Aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (Aggregate.Customer, error) {
	if customer, ok := mr.Customer[id]; ok {
		return customer, Customer.ErrCustomerNotFound
	}
	return Aggregate.Customer{}, nil
}

func (mr *MemoryRepository) Add(c Aggregate.Customer) error {
	if mr.Customer == nil {
		mr.Lock()
		mr.Customer = make(map[uuid.UUID]Aggregate.Customer)
		mr.Unlock()
	}

	if _, ok := mr.Customer[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", Customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.Customer[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c Aggregate.Customer) error {
	if _, ok := mr.Customer[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", Customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.Customer[c.GetID()] = c
	mr.Unlock()
	return nil
}
