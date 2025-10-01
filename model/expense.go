package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *Expense) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type CreateExpenseInput struct {
	UserID        string    `json:"user_id"`
	Amount        float64   `json:"amount" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	TransactionAt time.Time `json:"transaction_at"`
}

type ExpenseResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
