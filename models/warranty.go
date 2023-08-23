package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Warranty struct {
	ID              uuid.UUID `json:"id" db:"id" rw:"r"`
	TransactionTime time.Time `json:"transaction_time" db:"transaction_time"`
	ExpiryTime      time.Time `json:"expiry_time" db:"expiry_time"`
	BrandName       string    `json:"brand_name" db:"brand_name"`
	StoreName       string    `json:"store_name" db:"store_name"`
	Amount          int       `json:"amount" db:"amount"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Uid             uuid.UUID `json:"uid" db:"uid"`
}
