package ValueObject

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Amount   int
	from     uuid.UUID
	to       uuid.UUID
	createAt time.Time
}
