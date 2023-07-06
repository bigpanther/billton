package warranty

import "time"

type Warranty struct {
	ID              int
	TransactionDate time.Time
	Expiry          time.Time
	Brand           string
	Amount          int // 100/100
}
