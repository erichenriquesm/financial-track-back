package route

import (
	"financial-track/controller"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseRoutes(r *gin.RouterGroup) {
	expense := r.Group("/expenses")
	{
		expense.POST("/", controller.CreateExpense)
		expense.GET("/mensal-summary", controller.GetMensalSummary)
	}
}
