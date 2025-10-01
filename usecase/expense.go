package usecase

import (
	"errors"
	"financial-track/model"
	"financial-track/repository"

	"github.com/google/uuid"
)

type ExpenseUseCase struct {
	repo *repository.ExpenseRepository
}

func NewExpenseUseCase(repo *repository.ExpenseRepository) *ExpenseUseCase {
	return &ExpenseUseCase{repo: repo}
}

func (e *ExpenseUseCase) CreateExpense(input model.CreateExpenseInput) (model.Expense, error) {
	if input.Amount <= 0 {
		return model.Expense{}, errors.New("invalid amount")
	}
	if input.Description == "" {
		return model.Expense{}, errors.New("description cannot be empty")
	}

	input.TransactionAt = input.TransactionAt.UTC()

	userId, err := uuid.Parse(input.UserID)
	if err != nil {
		return model.Expense{}, errors.New("invalid user id")
	}

	expense := model.Expense{
		UserID:        userId,
		Amount:        input.Amount,
		Description:   input.Description,
		TransactionAt: input.TransactionAt,
	}

	if err := e.repo.Create(&expense); err != nil {
		return model.Expense{}, err
	}
	return expense, nil
}
