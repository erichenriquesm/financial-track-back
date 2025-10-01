package repository

import (
	"financial-track/database"
	"financial-track/model"
)

type ExpenseRepository struct{}

func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{}
}

func (r *ExpenseRepository) Create(expense *model.Expense) error {
	return database.DB.Create(expense).Error
}
