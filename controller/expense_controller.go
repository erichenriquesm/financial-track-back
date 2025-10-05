package controller

import (
	"financial-track/model"
	"financial-track/repository"
	"financial-track/usecase"
	"financial-track/utils"
	"strconv"
	"time"

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

func GetMensalSummary(c *gin.Context) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endDate := time.Now()

	page := 1
	pageSize := 15
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if ps := c.Query("perPage"); ps != "" {
		if v, err := strconv.Atoi(ps); err == nil && v > 0 {
			pageSize = v
		}
	}

	var body model.PaginationParams
	if err := c.ShouldBindJSON(&body); err == nil {
		if body.Page > 0 {
			page = body.Page
		}
		if body.PageSize > 0 {
			pageSize = body.PageSize
		}
	}

	paged, err := expenseUseCase.GetMensalSummary(startDate, endDate, page, pageSize)
	if err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	c.JSON(200, paged)
}
