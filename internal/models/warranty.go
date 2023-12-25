package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warranty struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	TransactionTime time.Time `json:"transaction_time"`
	ExpiryTime      time.Time `json:"expiry_time"`
	BrandName       string    `json:"brand_name" gorm:"size:255;not null"`
	StoreName       string    `json:"store_name" gorm:"size:255;not null"`
	Amount          int       `json:"amount"`
	UserID          uuid.UUID `json:"userid" gorm:"type:uuid;not null;index"`
	User            User
}
