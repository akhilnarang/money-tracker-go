package models

import (
	"github.com/google/uuid"
	"time"
)

type Expense struct {
	Id            uuid.UUID `json:"id" gorm:"primaryKey"`
	Description   string    `json:"description"`
	Amount        float32   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Created       time.Time `json:"created" gorm:"autoCreateTime"`
	LastUpdated   time.Time `json:"last_updated" gorm:"autoUpdateTime"`
}

type ExpenseUpdate struct {
	Description   string  `json:"description"`
	Amount        float32 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}
