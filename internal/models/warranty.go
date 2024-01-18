package models

import (
	"time"

	"github.com/google/uuid"
)

type Warranty struct {
	Model
	TransactionTime time.Time `json:"transaction_time"`
	ExpiryTime      time.Time `json:"expiry_time"`
	BrandName       string    `json:"brand_name" gorm:"size:255;not null"`
	StoreName       string    `json:"store_name" gorm:"size:255;not null"`
	Amount          int       `json:"amount"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User            User
}
