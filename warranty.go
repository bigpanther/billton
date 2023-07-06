package main

import "time"

type warranty struct {
	ID              int
	TransactionDate time.Time
	Expiry          time.Time
	Brand           string
	Amount          int // 100/100
}
