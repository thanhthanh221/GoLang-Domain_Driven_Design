package Customer

import (
	"GolangDDD/Aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was found in the reposotyry")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update teh customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (Aggregate.Customer, error)
	Add(Aggregate.Customer) error
	Update(Aggregate.Customer) error
}
