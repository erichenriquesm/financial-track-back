package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category string

const (
	Food           Category = "FOOD"
	Transportation Category = "TRANSPORTATION"
	Housing        Category = "HOUSING"
	Health         Category = "HEALTH"
	Education      Category = "EDUCATION"
	Entertainment  Category = "ENTERTAINMENT"
	Clothing       Category = "CLOTHING"
	Personal       Category = "PERSONAL"
	Finance        Category = "FINANCE"
	Others         Category = "OTHERS"
)

var validCategories = []Category{
	Food, Transportation, Housing, Health, Education,
	Entertainment, Clothing, Personal, Finance, Others,
}

func IsValidCategory(c Category) bool {
	for _, cat := range validCategories {
		if c == cat {
			return true
		}
	}
	return false
}

type Expense struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"index;index:idx_user_transaction_at,priority:1" json:"userId"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Category      Category  `gorm:"type:varchar(20)" json:"category"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	TransactionAt time.Time `gorm:"index;index:idx_user_transaction_at,priority:2,sort:desc" json:"transactionAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (u *Expense) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type CreateExpenseInput struct {
	UserID        string   `json:"userId"`
	Category      Category `gorm:"type:varchar(20)" json:"category" binding:"required"`
	Amount        float64  `json:"amount" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	TransactionAt JSONTime `json:"transactionAt" binding:"required"`
}

type ExpenseResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"userId"`
	Category      Category  `gorm:"type:varchar(20)" json:"category" binding:"required"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	TransactionAt time.Time `json:"transactionAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Summary struct {
	TotalAmount float64    `json:"total_amount"`
	Pagination  Pagination `json:"pagination"`
}

type PagedSummary struct {
	Amount      float64           `json:"amount"`
	Data        []ExpenseResponse `json:"data"`
	CurrentPage int               `json:"currentPage"`
	LastPage    int               `json:"lastPage"`
	TotalItems  int64             `json:"totalItems"`
	PerPage     int               `json:"perPage"`
}
