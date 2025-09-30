package route

import "github.com/gin-gonic/gin"

func RegisterHealthRoutes(r *gin.Engine) {
	healthcheck := r.Group("/")
	{
		healthcheck.GET("/ping", pingRoute)
	}
}

func pingRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
