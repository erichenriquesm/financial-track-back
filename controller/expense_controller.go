package controller

import (
	"financial-track/model"
	"financial-track/repository"
	"financial-track/usecase"
	"financial-track/utils"

	"github.com/gin-gonic/gin"
)

var expenseRepository *repository.ExpenseRepository = repository.NewExpenseRepository()
var expenseUseCase *usecase.ExpenseUseCase = usecase.NewExpenseUseCase(expenseRepository)

func CreateExpense(c *gin.Context) {
	var createExpenseInput model.CreateExpenseInput
	userId, err := c.Get("userId")

	if !err {
		c.JSON(400, gin.H{"errors": "User ID not found in context"})
		return
	}

	errs := utils.ValidateJSON(c, &createExpenseInput)
	if errs != nil {
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if !model.IsValidCategory(createExpenseInput.Category) {
		c.JSON(400, gin.H{"errors": "Invalid category"})
		return
	}

	createExpenseInput.UserID = userId.(string)

	expense, expenseErr := expenseUseCase.CreateExpense(createExpenseInput)

	if expenseErr != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	resp := model.ExpenseResponse{
		ID:            expense.ID,
		UserID:        expense.UserID,
		Category:      expense.Category,
		Amount:        expense.Amount,
		Description:   expense.Description,
		TransactionAt: expense.TransactionAt,
		CreatedAt:     expense.CreatedAt,
		UpdatedAt:     expense.UpdatedAt,
	}

	c.JSON(201, gin.H{"message": "Expense created successfully", "expense": resp})
}

func GetExpenses(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GetExpenses endpoint hit"})
}
