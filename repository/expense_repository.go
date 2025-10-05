package repository

import (
	"financial-track/database"
	"financial-track/model"
	"time"
)

type ExpenseRepository struct{}

func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{}
}

func (r *ExpenseRepository) Create(expense *model.Expense) error {
	return database.DB.Create(expense).Error
}

func (r *ExpenseRepository) GetSummary(startDate, endDate time.Time, page, pageSize int) (model.PagedSummary, error) {
	var summary model.Summary

	if err := database.DB.Model(&model.Expense{}).
		Where("transaction_at BETWEEN ? AND ?", startDate, endDate).
		Select("COALESCE(SUM(amount),0)").Scan(&summary.TotalAmount).Error; err != nil {
	}

	var totalItems int64
	if err := database.DB.Model(&model.Expense{}).
		Where("transaction_at BETWEEN ? AND ?", startDate, endDate).
		Count(&totalItems).Error; err != nil {
		return model.PagedSummary{}, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 15
	}
	offset := (page - 1) * pageSize

	var expensesDB []model.Expense
	if err := database.DB.
		Where("transaction_at BETWEEN ? AND ?", startDate, endDate).
		Order("transaction_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&expensesDB).Error; err != nil {
		return model.PagedSummary{}, err
	}

	totalPages := 0
	if pageSize > 0 {
		totalPages = int((totalItems + int64(pageSize) - 1) / int64(pageSize))
	}

	expenses := make([]model.ExpenseResponse, 0, len(expensesDB))
	for _, e := range expensesDB {
		expenses = append(expenses, model.ExpenseResponse{
			ID:            e.ID,
			UserID:        e.UserID,
			Category:      e.Category,
			Amount:        e.Amount,
			Description:   e.Description,
			TransactionAt: e.TransactionAt,
			CreatedAt:     e.CreatedAt,
			UpdatedAt:     e.UpdatedAt,
		})
	}

	return model.PagedSummary{
		Amount:      summary.TotalAmount,
		Data:        expenses,
		CurrentPage: page,
		LastPage:    totalPages,
		TotalItems:  totalItems,
		PerPage:     pageSize,
	}, nil
}
