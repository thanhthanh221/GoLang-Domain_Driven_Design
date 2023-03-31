package Memory

import (
	"GolangDDD/Aggregate"
	"GolangDDD/Domain/Customer"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name         string
		id           uuid.UUID
		exepectedErr error
	}

	cust, err := Aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		Customer: map[uuid.UUID]Aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:         "not customer by id",
			id:           uuid.MustParse("c72837fc-81cf-4ed0-ae67-7f120b3ab573"),
			exepectedErr: Customer.ErrCustomerNotFound,
		},
		{
			name:         "customer",
			id:           id,
			exepectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.exepectedErr) {
				t.Errorf("expected error %v, got %v", tc.exepectedErr, err)
			}
		})
	}
}
